package postgres

import (
	"AirLineTicketTracker/config"
	"AirLineTicketTracker/internal/entities"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(config *config.Config) (*Storage, error) {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable",
		config.Database.Username, config.Database.Password,
		config.Database.Host, config.Database.Port,
		config.Database.DbName)
	fmt.Println("Connecting to PostgreSQL...", connString)

	conn, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Println("Error connecting to database")
		return nil, err
	}

	return &Storage{conn}, nil
}

func (s *Storage) StoreUser(user entities.User) error {
	users := make([]UsersSchema, 0)
	err := s.db.Select(&users, "SELECT * FROM users WHERE chat_id=$1 AND username=$2",
		user.ChatId, user.Username)
	if err != nil {
		return err
	}

	if len(users) > 0 {
		return UserAlreadyExistsError
	}

	var id uint64
	err = s.db.Get(&id, "INSERT INTO users (chat_id, username) VALUES($1, $2) RETURNING id", user.ChatId, user.Username)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) StoreUserFlight(chatId int64, flight *entities.Flight) error {
	users := make([]UsersSchema, 0)
	// Проверяем есть ли пользователь в базе данных
	err := s.db.Select(&users, "SELECT * FROM users WHERE chat_id=$1", chatId)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return UserNotFoundError
	}

	flights := make([]FlightsSchema, 0)
	// Проверяем отслеживается ли уже этот полет
	err = s.db.Select(&flights,
		"SELECT * FROM flights WHERE origin_iata=$1 AND destination_iata=$2",
		flight.OriginIATA, flight.DestinationIATA)
	if err != nil {
		return err
	}

	if len(flights) > 0 {
		return FlightAlreadyStored
	}

	var id uint64
	err = s.db.Get(&id, `INSERT INTO flights
    			(origin_iata, origin, destination_iata, destination,
    			 price, departure_at, user_id) 
				VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		flight.OriginIATA, flight.Origin, flight.DestinationIATA, flight.Destination,
		flight.Price, flight.DepartureAt, users[0].Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetUserFlight(chatId int64) ([]entities.Flight, error) {
	flights := make([]FlightsSchema, 0)
	err := s.db.Select(&flights, `
		SELECT 
			f.id,
			f.origin_iata,
			f.origin,
			f.destination_iata,
			f.destination,
			f.price,
			f.departure_at,
			f.user_id
		FROM 
			flights f
		JOIN 
			users u 
		ON 
			f.user_id = u.id
		WHERE 
			u.chat_id = $1;
	`, chatId)

	if err != nil {
		return nil, err
	}

	// мапим данные из структуры схемы в структуру flight
	res := make([]entities.Flight, len(flights))
	for i, flight := range flights {
		res[i] = *flight.ToFlight()
	}

	return res, nil
}

func (s *Storage) GetFlightId(chatId int64, flight *entities.Flight) (int64, error) {
	flights := make([]int64, 0)

	err := s.db.Select(&flights, `
		SELECT 
			f.id
		FROM 
			flights f
		JOIN 
			users u 
		ON 
			f.user_id = u.id
		WHERE 
			u.chat_id = $1 and f.origin_iata = $2 and f.destination_iata = $3;
	`, chatId, flight.OriginIATA, flight.DestinationIATA)

	if err != nil {
		return 0, err
	}

	return flights[0], nil
}

func (s *Storage) GetFlightById(id int64) (*entities.Flight, error) {
	flight := new(FlightsSchema)
	err := s.db.Get(flight, "SELECT * FROM flights WHERE id = $1", id)

	if err != nil {
		return nil, err
	}
	//f := flight
	return flight.ToFlight(), nil
}

func (s *Storage) ChangeFlightPrice(id int64, price int) error {
	var n *int64
	err := s.db.Get(n, "UPDATE flights SET price = $1 WHERE id=$2", price, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetAllFlights() ([]entities.Flight, error) {
	flights := make([]FlightsSchema, 0)
	err := s.db.Select(&flights, "SELECT * FROM flights")
	if err != nil {
		return nil, err
	}

	res := make([]entities.Flight, len(flights))
	for i, flight := range flights {
		res[i] = *flight.ToFlight()
	}

	return res, nil
}

func (s *Storage) GetUsers() ([]entities.User, error) {
	// Выполнение SQL-запроса
	rows, err := s.db.Query(`
		SELECT 
			u.username,
			u.chat_id,
			f.origin_iata,
			f.origin,
			f.destination_iata,
			f.destination,
			f.price,
			f.departure_at
		FROM 
			users u
		LEFT JOIN 
			flights f
		ON 
			u.id = f.user_id;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Мапинг данных в структуры
	usersMap := make(map[string]*entities.User)

	for rows.Next() {
		var username string
		var chatId int64
		var originIATA, origin, destinationIATA, destination sql.NullString
		var price sql.NullInt64
		var departureAt sql.NullTime

		// Чтение строки результата
		err := rows.Scan(&username, &chatId, &originIATA, &origin, &destinationIATA, &destination, &price, &departureAt)
		if err != nil {
			panic(err)
		}

		// Поиск или создание пользователя
		user, exists := usersMap[username]
		if !exists {
			user = &entities.User{
				Username: username,
				ChatId:   chatId,
				Flights:  []entities.Flight{},
			}
			usersMap[username] = user
		}

		// Добавление перелета (если он есть)
		if originIATA.Valid && destinationIATA.Valid {
			flight := entities.Flight{
				OriginIATA:      originIATA.String,
				Origin:          origin.String,
				DestinationIATA: destinationIATA.String,
				Destination:     destination.String,
				Price:           int(price.Int64),
				DepartureAt:     departureAt.Time,
			}
			user.Flights = append(user.Flights, flight)
		}
	}

	// Конвертация карты пользователей в массив (если нужно)
	var users []entities.User
	for _, user := range usersMap {
		users = append(users, *user)
	}

	return users, nil
}
