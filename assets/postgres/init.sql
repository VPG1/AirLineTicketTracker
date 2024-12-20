CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    chat_id INT NOT NULL,
    username VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS flights
(
    id SERIAL PRIMARY KEY,
    origin_iata VARCHAR(3) NOT NULL,
    origin VARCHAR(20) NOT NULL,
    destination_iata VARCHAR(3) NOT NULL,
    destination VARCHAR(20) NOT NULL,
    price INT NOT NULL,
    departure_at TIMESTAMP WITH TIME ZONE NOT NULL,
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE
);
