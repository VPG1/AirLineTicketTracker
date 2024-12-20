package telegram

const (
	startResponse = `
	Привет! 👋 Я бот для отслеживания авиабилетов.
	Я помогу тебе найти лучшие цены на рейсы и уведомлю, когда они снизятся.
	
	Вот что я могу:
	- ✈️ Отслеживать цены на билеты.
	- 🔔 Уведомлять о снижении цен.
	- 🕒 Присылать регулярные отчёты.
	
	Напиши /help, чтобы узнать подробнее о командах
`

	helpResponse = `
	Вот список доступных команд:  
	
	- /track <город вылета> <город назначения>  
	  ➡️ Добавить рейс для отслеживания.  
	
	- /list  
	  📋 Показать все рейсы, которые ты отслеживаешь.  
	
	- /remove <номер рейса>  
	  🗑 Удалить рейс из списка.  
	
	- /notifications <on|off>  
	  🔔 Включить или выключить уведомления.  
	
	- /settings  
	  ⚙️ Настройки уведомлений.  
	
	Если у тебя есть вопросы или предложения, напиши нам!
`
	trackResponse = `
	Рейс успешно добавлен для отслеживания! ✈️  
	📍 Откуда: %s 
	📍 Куда: %s 	
	🗓 Дата: %s
	💰 Текущая цена: %v$
	Мы уведомим тебя, если цена изменится.
`
	incorrectSearchPhraseResponse = `
	Упс! 😔 Не удалось добавить рейс.  
	Проверь формат команды:  
	/track <город вылета> <город назначения>
`
	userNotInSystem = `
	Упс! 😔Вы не зарегистрированы в системе.  
	Пожалуйста введите команд /start
`

	userAlreadyTrackFlight = `
	Вы уже отслеживаете рейс.
	📍 Откуда: %s 
	📍 Куда: %s 
	🗓 Дата: %s
	💰 Текущая цена: %v$
	Мы уведомим тебя, если цена изменится.
`
	listResponse = `
	Вот твои отслеживаемые рейсы:  

	%s
	Напиши /remove <номер рейса>, чтобы удалить рейс из списка.
`
	NoTrackedFlights = `
	Ты ещё не добавил ни одного рейса для отслеживания. 😕
`

	Notification = `
	📉 Отличные новости! Цена на рейс снизилась!  

	📍 Откуда: %s 
	📍 Куда: %s
	🗓 Дата: %s 
	💰 Новая цена: %v
	
	Поторопись, предложение может быстро исчезнуть!
`
)
