package data

type Skill struct {
	Name  string
	Level int
}

type Project struct {
	Name         string
	Description  string
	Technologies []string
	GitHubURL    string
}

type Experience struct {
	Position    string
	Company     string
	Period      string
	Description string
}

type Education struct {
	Degree      string
	Institution string
	Year        string
}

type Contact struct {
	Email    string
	Phone    string
	Location string
	GitHub   string
	LinkedIn string
}

var ResumeData = struct {
	Name       string
	Title      string
	Summary    string
	Skills     []Skill
	Projects   []Project
	Experience []Experience
	Education  []Education
	Contact    Contact
}{
	Name:  "Иван Иванов",
	Title: "Go Разработчик",
	Summary: "Опытный Go-разработчик с 3+ годами опыта в создании высоконагруженных распределенных систем. " +
		"Специализируюсь на микросервисной архитектуре, REST API и облачных технологиях.",

	Skills: []Skill{
		{"Go", 90},
		{"Docker", 85},
		{"Kubernetes", 80},
		{"PostgreSQL", 85},
		{"Redis", 75},
		{"gRPC", 80},
		{"Git", 90},
		{"Linux", 85},
	},

	Projects: []Project{
		{
			Name:         "Микросервисная платформа",
			Description:  "Разработка платформы для обработки 1M+ запросов в день",
			Technologies: []string{"Go", "Docker", "Kubernetes", "gRPC"},
			GitHubURL:    "https://github.com/username/microplatform",
		},
		{
			Name:         "REST API сервис",
			Description:  "Высокопроизводительный API с JWT аутентификацией",
			Technologies: []string{"Go", "PostgreSQL", "Redis", "JWT"},
			GitHubURL:    "https://github.com/username/api-service",
		},
	},

	Experience: []Experience{
		{
			Position:    "Senior Go Developer",
			Company:     "TechCorp Inc.",
			Period:      "2022 - настоящее время",
			Description: "Разработка высоконагруженных микросервисов, оптимизация производительности",
		},
		{
			Position:    "Go Developer",
			Company:     "StartUp Solutions",
			Period:      "2020 - 2022",
			Description: "Создание backend-систем с нуля, работа с базами данных",
		},
	},

	Education: []Education{
		{
			Degree:      "Магистр компьютерных наук",
			Institution: "Технический университет",
			Year:        "2020",
		},
		{
			Degree:      "Бакалавр программной инженерии",
			Institution: "Университет информационных технологий",
			Year:        "2018",
		},
	},

	Contact: Contact{
		Email:    "ivan.ivanov@email.com",
		Phone:    "+7 (999) 123-45-67",
		Location: "Москва, Россия",
		GitHub:   "https://github.com/username",
		LinkedIn: "https://linkedin.com/in/username",
	},
}
