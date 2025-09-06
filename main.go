package main

import (
	"fmt"
	"image/color"
	"net/url"
	"resume-app/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

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

func createMainTab() *container.TabItem {
	title := canvas.NewText(data.ResumeData.Name, color.NRGBA{R: 0, G: 100, B: 200, A: 255})
	title.TextSize = 24
	title.TextStyle.Bold = true

	subtitle := canvas.NewText(data.ResumeData.Title, color.NRGBA{R: 100, G: 100, B: 100, A: 255})
	subtitle.TextSize = 18

	summary := widget.NewLabel(data.ResumeData.Summary)
	summary.Wrapping = fyne.TextWrapWord

	contactInfo := widget.NewLabel(fmt.Sprintf(
		"📧 %s\n📞 %s\n📍 %s",
		data.ResumeData.Contact.Email,
		data.ResumeData.Contact.Phone,
		data.ResumeData.Contact.Location,
	))

	// Статистика навыков
	statsLabel := widget.NewLabel("📊 Статистика навыков:")
	statsLabel.TextStyle.Bold = true

	totalSkills := len(data.ResumeData.Skills)
	advancedSkills := 0
	for _, skill := range data.ResumeData.Skills {
		if skill.Level >= 80 {
			advancedSkills++
		}
	}

	statsText := widget.NewLabel(fmt.Sprintf(
		"• Всего навыков: %d\n• Продвинутых (80%%+): %d\n• Средний уровень: %.1f%%",
		totalSkills,
		advancedSkills,
		calculateAverageLevel(data.ResumeData.Skills),
	))

	mainContent := container.NewVBox(
		title,
		subtitle,
		layout.NewSpacer(),
		summary,
		layout.NewSpacer(),
		contactInfo,
		layout.NewSpacer(),
		statsLabel,
		statsText,
	)

	return container.NewTabItem("Главная", mainContent)
}

func createSkillsTab() *container.TabItem {
	skillsContainer := container.NewVBox()
	skillsContainer.Add(widget.NewLabel("💻 Технические навыки:"))
	skillsContainer.Add(widget.NewSeparator())

	for _, skill := range data.ResumeData.Skills {
		skillRow := container.NewVBox(
			container.NewHBox(
				widget.NewLabel(fmt.Sprintf("• %s", skill.Name)),
				layout.NewSpacer(),
				widget.NewLabel(fmt.Sprintf("%d%%", skill.Level)),
			),
			widget.NewProgressBar().SetValue(float64(skill.Level)/100),
		)
		skillsContainer.Add(skillRow)
		skillsContainer.Add(layout.NewSpacer())
	}

	scroll := container.NewScroll(skillsContainer)
	return container.NewTabItem("Навыки", scroll)
}

func createProjectsTab() *container.TabItem {
	projectsContainer := container.NewVBox()

	for i, project := range data.ResumeData.Projects {
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

	for _, exp := range data.ResumeData.Experience {
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
				widget.NewLabel("🏆 Достижения:"),
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

	for _, edu := range data.ResumeData.Education {
		eduCard := widget.NewCard(
			fmt.Sprintf("🎓 %s", edu.Degree),
			edu.Institution,
			container.NewVBox(
				widget.NewLabel(fmt.Sprintf("📅 Год: %s", edu.Year)),
				widget.NewLabel(fmt.Sprintf("📚 Специализация: %s", edu.Specialization)),
			),
		)
		eduContainer.Add(eduCard)
		eduContainer.Add(layout.NewSpacer())
	}

	scroll := container.NewScroll(eduContainer)
	return container.NewTabItem("Образование", scroll)
}

func createCertificatesTab() *container.TabItem {
	certContainer := container.NewVBox()

	for _, cert := range data.ResumeData.Certificates {
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
	contact := data.ResumeData.Contact

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
		widget.NewLabel("• Готов к удаленной работе\n• Рассматриваю офис в Москве\n• Полная занятость\n• Испытательный срок 3 месяца"),
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

func calculateAverageLevel(skills []data.Skill) float64 {
	total := 0
	for _, skill := range skills {
		total += skill.Level
	}
	return float64(total) / float64(len(skills))
}
