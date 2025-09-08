package main

import (
	"fmt"
	"image/color"
	"log"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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
	PhotoPath    string
}{
	Name:      "Демид Пьянков",
	Title:     "Начинающий Go разработчик",
	PhotoPath: "photo.jpg",
	Summary: "Занимаюсь Go полгода в основном фронтендом, если надо могу и в бекэнд.\n " +
		"Умею читать код на Go и Pythone.\n " +
		"Готов учиться новому, учусь в быстром темпе.\n " +
		"В планах поступить на заочку в университет.\n ",

	Skills: []Skill{
		{"Go", 50},
		{"Python", 50},
		{"Git", 50},
		{"Настойчивость", 100},
		{"Стремления", 100},
		{"Навык общения", 100},
		{"Фокус на цели", 100},
		{"Умение работать с Ai", 100},
	},

	Projects: []Project{
		{
			Name:         "Maincraft-launcher",
			Description:  "Более современный и удобный способ",
			Technologies: []string{"Go", "С", "fyne", "javascript"},
		},
	},

	Experience: []Experience{
		{
			Position:    "Ассесор",
			Company:     "Яндекс",
			Period:      "2023 - 2024",
			Description: "фильтрация и оценка сайтов",
		},
	},

	Education: []Education{
		{
			Degree: "Среднее неполное (11 классов)",
		},
	},

	Certificates: []Certificate{
		{
			Name:   "PRO GO. Основы программирования",
			Issuer: "Stepick",
			Year:   "2025",
		},
	},

	Contact: Contact{
		Email:    "just.wwwerty@gmail.com",
		Phone:    "+7 915 278 97 37",
		Location: "Астрахань, Россия",
	},
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Резюме - Демид Пьянков")
	myWindow.Resize(fyne.NewSize(1000, 700))

	tabs := container.NewAppTabs(
		createMainTab(),
		createSkillsTab(),
		createProjectsTab(),
		createExperienceTab(),
		createEducationTab(),
		createCertificatesTab(),
		createContactTab(),
	)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func loadImage(path string) *canvas.Image {
	log.Printf("Пытаемся загрузить: %s", path)

	// Просто открываем файл по указанному пути
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Ошибка открытия файла: %v", err)
		log.Printf("Убедитесь что файл photo.jpg находится в той же папке что и main.go")
		return createDefaultAvatar()
	}
	defer file.Close()

	img := canvas.NewImageFromReader(file, "photo.jpg")
	img.SetMinSize(fyne.NewSize(150, 150))
	img.FillMode = canvas.ImageFillContain
	log.Printf("Изображение успешно загружено")
	return img
}

func createDefaultAvatar() *canvas.Image {
	// Используем иконку из темы
	avatar := canvas.NewImageFromResource(theme.AccountIcon())
	avatar.SetMinSize(fyne.NewSize(150, 150))
	avatar.FillMode = canvas.ImageFillContain
	return avatar
}

func createMainTab() *container.TabItem {
	title := canvas.NewText(ResumeData.Name, color.NRGBA{R: 0, G: 100, B: 200, A: 255})
	title.TextSize = 40
	title.TextStyle.Bold = true

	subtitle := canvas.NewText(ResumeData.Title, color.NRGBA{R: 100, G: 100, B: 100, A: 255})
	subtitle.TextSize = 25

	summary := widget.NewLabel(ResumeData.Summary)
	summary.Wrapping = fyne.TextWrapWord

	photo := loadImage(ResumeData.PhotoPath)

	summary = widget.NewLabel(ResumeData.Summary)
	summary.Wrapping = fyne.TextWrapWord

	contactInfo := widget.NewLabel(fmt.Sprintf(
		"📧 %s\n📞 %s\n📍 %s",
		ResumeData.Contact.Email,
		ResumeData.Contact.Phone,
		ResumeData.Contact.Location,
	))
	header := container.NewHBox(
		photo,
		container.NewVBox(
			title,
			subtitle,
			layout.NewSpacer(),
			summary,
		),
	)

	mainContent := container.NewVBox(
		header,
		layout.NewSpacer(),
		contactInfo,
	)

	return container.NewTabItem("Главная", mainContent)

}

func createSkillsTab() *container.TabItem {
	skillsContainer := container.NewVBox()
	skillsContainer.Add(widget.NewLabel("💻 Технические навыки:"))
	skillsContainer.Add(widget.NewSeparator())

	for _, skill := range ResumeData.Skills {
		// Создаем прогресс-бар и устанавливаем значение
		progressBar := widget.NewProgressBar()
		progressBar.SetValue(float64(skill.Level) / 100) // Отдельная операция

		skillRow := container.NewVBox(
			container.NewHBox(
				widget.NewLabel(fmt.Sprintf("• %s", skill.Name)),
				layout.NewSpacer(),
				widget.NewLabel(fmt.Sprintf("%d%%", skill.Level)),
			),
			progressBar, // Используем уже настроенный прогресс-бар
		)
		skillsContainer.Add(skillRow)
		skillsContainer.Add(layout.NewSpacer())
	}

	scroll := container.NewScroll(skillsContainer)
	return container.NewTabItem("Навыки", scroll)
}

func createProjectsTab() *container.TabItem {
	projectsContainer := container.NewVBox()

	for i, project := range ResumeData.Projects {
		projectCard := widget.NewCard(
			fmt.Sprintf("🚀 Проект %d: %s", i+1, project.Name),
			project.Description,
			container.NewVBox(
				widget.NewLabel("🛠️ Технологии:"),
				widget.NewLabel("  "+joinStrings(project.Technologies)),
			),
		)
		projectsContainer.Add(projectCard)
		projectsContainer.Add(layout.NewSpacer())
	}

	scroll := container.NewScroll(projectsContainer)
	return container.NewTabItem("Проекты", scroll)
}

func createExperienceTab() *container.TabItem {
	expContainer := container.NewVBox()

	for _, exp := range ResumeData.Experience {
		achievements := ""
		for _, achievement := range exp.Achievements {
			achievements += fmt.Sprintf("• %s\n", achievement)
		}

		expCard := widget.NewCard(
			fmt.Sprintf("🏢 %s", exp.Company),
			fmt.Sprintf("💼 %s | 📅 %s", exp.Position, exp.Period),
			container.NewVBox(
				widget.NewLabel(exp.Description),
				widget.NewSeparator(),
				widget.NewLabel(achievements),
			),
		)
		expContainer.Add(expCard)
		expContainer.Add(layout.NewSpacer())
	}

	scroll := container.NewScroll(expContainer)
	return container.NewTabItem("Опыт работы", scroll)
}

func createEducationTab() *container.TabItem {
	eduContainer := container.NewVBox()

	for _, edu := range ResumeData.Education {
		eduCard := widget.NewCard(
			fmt.Sprintf("🎓 %s", edu.Degree),
			edu.Institution,
			container.NewVBox(),
		)
		eduContainer.Add(eduCard)
		eduContainer.Add(layout.NewSpacer())
	}

	scroll := container.NewScroll(eduContainer)
	return container.NewTabItem("Образование", scroll)
}

func createCertificatesTab() *container.TabItem {
	certContainer := container.NewVBox()

	for _, cert := range ResumeData.Certificates {
		certCard := widget.NewCard(
			fmt.Sprintf("📜 %s", cert.Name),
			cert.Issuer,
			widget.NewLabel(fmt.Sprintf("📅 Год получения: %s", cert.Year)),
		)
		certContainer.Add(certCard)
		certContainer.Add(layout.NewSpacer())
	}

	scroll := container.NewScroll(certContainer)
	return container.NewTabItem("Сертификаты", scroll)
}

func createContactTab() *container.TabItem {
	contact := ResumeData.Contact

	emailBtn := widget.NewButton("📧 Email: "+contact.Email, func() {
		openURL("mailto:" + contact.Email)
	})

	phoneBtn := widget.NewButton("📞 Телефон: "+contact.Phone, func() {
		dialog.ShowInformation("Контактная информация",
			fmt.Sprintf("Телефон: %s\nВы можете позвонить или отправить SMS", contact.Phone),
			fyne.CurrentApp().Driver().AllWindows()[0])
	})

	locationBtn := widget.NewButton("📍 Местоположение: "+contact.Location, func() {
		dialog.ShowInformation("Местоположение",
			fmt.Sprintf("Город: %s\nГотов к релокации или удаленной работе", contact.Location),
			fyne.CurrentApp().Driver().AllWindows()[0])
	})

	// Информация о готовности к работе
	availabilityInfo := widget.NewCard(
		"✅ Доступность",
		"Трудовые условия",
		widget.NewLabel("• Готов к удаленной работе\n• Рассматриваю релокацию\n• Полная занятость"),
	)

	contactForm := container.NewVBox(
		widget.NewLabel("📞 Контактная информация:"),
		emailBtn,
		phoneBtn,
		locationBtn,
		layout.NewSpacer(),
		availabilityInfo,
	)

	return container.NewTabItem("Контакты", contactForm)
}

func openURL(urlStr string) {
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Printf("Ошибка открытия URL: %v\n", err)
		return
	}
	err = fyne.CurrentApp().OpenURL(u)
	if err != nil {
		fmt.Printf("Ошибка открытия: %v\n", err)
	}
}

func joinStrings(strs []string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += ", "
		}
		result += s
	}
	return result
}
