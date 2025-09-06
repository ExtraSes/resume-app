package data // Важно: package data, не main!

// Структура для навыков
type Skill struct {
	Name  string
	Level int
}

// Структура для проектов
type Project struct {
	Name         string
	Description  string
	Technologies []string
}

// Структура для опыта работы
type Experience struct {
	Position     string
	Company      string
	Period       string
	Description  string
	Achievements []string
}

// Структура для образования
type Education struct {
	Degree         string
	Institution    string
	Year           string
	Specialization string
}

// Структура для контактной информации
type Contact struct {
	Email    string
	Phone    string
	Location string
}

// Структура для сертификатов
type Certificate struct {
	Name   string
	Issuer string
	Year   string
}

// Экспортируемая переменная (с большой буквы!)
var ResumeData = struct {
	Name         string
	Title        string
	Summary      string
	Skills       []Skill
	Projects     []Project
	Experience   []Experience
	Education    []Education
	Certificates []Certificate
	Contact      Contact
}{
	Name:  "Иван Иванов",
	Title: "Go Разработчик",

	Summary: "Опытный Go-разработчик с 3+ годами опыта в создании высоконагруженных распределенных систем. " +
		"Специализируюсь на микросервисной архитектуре, REST API и облачных технологиях. " +
		"Сильные стороны: оптимизация производительности, проектирование масштабируемых систем, код ревью.",

	Skills: []Skill{
		{"Go", 90},
		{"Docker", 85},
		{"Kubernetes", 80},
		{"PostgreSQL", 85},
		{"Redis", 75},
		{"gRPC", 80},
		{"REST API", 88},
		{"Git", 90},
		{"Linux", 85},
		{"CI/CD", 75},
	},

	Projects: []Project{
		{
			Name:         "Микросервисная платформа",
			Description:  "Разработка высоконагруженной платформы для обработки более 1 миллиона запросов в день",
			Technologies: []string{"Go", "Docker", "Kubernetes", "gRPC", "PostgreSQL"},
		},
		{
			Name:         "REST API сервис",
			Description:  "Создание высокопроизводительного API сервиса с JWT аутентификацией и кэшированием",
			Technologies: []string{"Go", "PostgreSQL", "Redis", "JWT", "Docker"},
		},
	},

	Experience: []Experience{
		{
			Position:    "Senior Go Developer",
			Company:     "TechCorp Inc.",
			Period:      "2022 - настоящее время",
			Description: "Разработка и поддержка высоконагруженных микросервисов",
			Achievements: []string{
				"Увеличил производительность системы на 40% через оптимизацию кода",
				"Внедрил систему мониторинга и алертинга",
			},
		},
	},

	Education: []Education{
		{
			Degree:         "Магистр компьютерных наук",
			Institution:    "Технический университет",
			Year:           "2020",
			Specialization: "Распределенные системы",
		},
	},

	Certificates: []Certificate{
		{
			Name:   "Go Advanced Programming",
			Issuer: "Coursera",
			Year:   "2023",
		},
	},

	Contact: Contact{
		Email:    "ivan.ivanov@email.com",
		Phone:    "+7 (999) 123-45-67",
		Location: "Москва, Россия",
	},
}
