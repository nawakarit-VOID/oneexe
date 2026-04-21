// Copyright (c) 2026 Nawakarit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"embed"
	_ "embed"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type AppConfig struct {
	Name    string
	AppID   string
	Version string

	//exe
	CompanyName string
	Fileversion string
	Years       string
	License     string
}

// ============================================================================
// ฟังชั้น gen template
// ============================================================================
func generateFile(tmplPath, outputPath string, data AppConfig) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	//projectPath
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}

// ============================================================================
// ฟังชั้น build Scriptbuild EXE
// ============================================================================
func buildexe(projectPath string, output *widget.Entry) {

	commands := [][]string{
		{"gnome-terminal", "--", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildexe.sh && ./buildexe.sh; exec bash"},
		{"x-terminal-emulator", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildexe.sh && ./buildexe.sh; exec bash"},
		{"konsole", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildexe.sh && ./buildexe.sh; exec bash"},
		{"xfce4-terminal", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildexe.sh && ./buildexe.sh; exec bash"},
	}

	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		err := cmd.Start()
		if err == nil {
			output.SetText("🚀 opened terminal: " + c[0])
			return
		}
	}

	output.SetText("❌ no terminal found")
}

// ============================================================================
// ฟังชั้น build Icons
// ============================================================================
func runScriptbuildIcons(projectPath string, output *widget.Entry) {

	commands := [][]string{ //ใช้ imagemagick
		{"gnome-terminal", "--", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildicons.sh && ./buildicons.sh; exec bash"},
		{"x-terminal-emulator", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildicons.sh && ./buildicons.sh; exec bash"},
		{"konsole", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildicons.sh && ./buildicons.sh; exec bash"},
		{"xfce4-terminal", "-e", "bash", "-c", "cd '" + projectPath + "' && chmod +x buildicons.sh && ./buildicons.sh; exec bash"},
	}

	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		err := cmd.Start()
		if err == nil {
			output.SetText("🚀 opened terminal: " + c[0])
			return
		}
	}

	output.SetText("❌ no terminal found")
}

// โหลด icon
func loadIcon(size int) fyne.Resource {
	var file string

	switch {
	case size >= 512:
		file = "icons/icon-512.png" ///ที่อยู่
	case size >= 256:
		file = "icons/icon-256.png"
	case size >= 128:
		file = "icons/icon-128.png"
	default:
		file = "icons/icon-64.png"
	}

	data, _ := iconFS.ReadFile(file)
	return fyne.NewStaticResource(file, data)
}

//go:embed icons/*
var iconFS embed.FS

func main() {

	a := app.NewWithID("com.nawakarit.oneexe")
	icons := loadIcon(64) //เอา data มาใช้
	a.SetIcon(icons)
	w := a.NewWindow("oneexe")
	w.SetIcon(icons)

	// inputs
	name := widget.NewEntry()
	name.SetText("")
	name.SetPlaceHolder("***App Name - ชื่อโปรแกรม-แอพ ")

	appID := widget.NewEntry()
	appID.SetText("com.nawakarit.pomodoro")
	appID.SetPlaceHolder("*com.example.app - แอพไอดี")

	version := widget.NewEntry()
	version.SetText("5.5.5")
	version.SetPlaceHolder("*ใ่ส่เวอร์ชัน เช่น 1.0.0")

	//exe
	companyName := widget.NewEntry()
	companyName.SetText("Nawakarit")
	companyName.SetPlaceHolder("*ชื่อบริษัท")

	fileversion := widget.NewEntry()
	fileversion.SetText("1,1,1,1")
	fileversion.SetPlaceHolder("*version (exe) เช่น 1,1,1,1")

	years := widget.NewLabel("")
	month := widget.NewLabel("")
	days := widget.NewLabel("")

	license := widget.NewEntry()
	license.SetText("GNU General Public License v3.0")
	license.SetPlaceHolder("*ใส่ประเภท license *ถ้าต้องการ")

	// log box
	logBox := widget.NewMultiLineEntry()
	logBox.SetPlaceHolder("Logs will appear here...")
	logBox.Wrapping = fyne.TextWrapWord

	// ============================================================================
	// เลือกแฟ้มเป้าหมาย
	// ============================================================================
	// 🔹 เลือก folder
	projectPath := ""

	selectBtn := widget.NewButton("1 - Select Project Folder", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil {
				return
			}

			projectPath = uri.Path()
			logBox.SetText("📁 Selected: " + projectPath)
		}, w)
	})

	// ============================================================================
	// Generate scrip Icons Btn
	// ============================================================================
	// 🔧 Generate
	genscripiconsBtn := widget.NewButton("2 - Generate scrip Icons", func() {

		if projectPath == "" {
			logBox.SetText("❌ Please select project folder")
			return
		}
		cfg := AppConfig{}

		generateFile("templates/buildicons.tmpl",
			filepath.Join(projectPath, "buildicons.sh"), cfg) //เอา scrip build ออกมาไว้นอกแฟ้ม flatpak

		logBox.SetText("✅ Generated File - - buildicons - -")
	})

	// ============================================================================
	// ปุ่ม Build Icons **ใช้ imagemagick
	// ============================================================================
	buildIconsBtn := widget.NewButton("3 - Run Build Icons", func() {

		if projectPath == "" {
			logBox.SetText("❌ select folder first")
			return
		}

		//  run script
		go runScriptbuildIcons(projectPath, logBox)

		logBox.SetText("🚀 Build started in terminal...")
	})

	// ============================================================================
	// Generate scrip Btn
	// ============================================================================
	// 🔧 Generate
	genscripexeBtn := widget.NewButton("Generate scrip EXE", func() {

		if projectPath == "" {
			logBox.SetText("❌ Please select project folder")
			return
		}

		cfg := AppConfig{
			Name:    name.Text,
			AppID:   appID.Text,
			Version: version.Text,

			//exe
			CompanyName: companyName.Text,
			Fileversion: fileversion.Text,
			Years:       years.Text,
			License:     license.Text,
		}

		flatpakPath := projectPath + "/" + "flatpak"
		os.MkdirAll(flatpakPath, 0755)

		generateFile("templates/app.rc.tmpl",
			filepath.Join(projectPath, "app.rc"), cfg) //เอา scrip build ออกมาไว้นอกแฟ้ม

		generateFile("templates/buildexe.tmpl",
			filepath.Join(projectPath, "buildexe.sh"), cfg)

		generateFile("templates/FyneApp.toml.tmpl",
			filepath.Join(projectPath, "FyneApp.toml"), cfg)

		//now := time.Now()
		//date.SetText(now.Format("2006-01-02"))
		//timeEntry.SetText(now.Format("15:04"))
		//years.SetText(now.Format("2006"))

		logBox.SetText("✅ Generated scrip exe")
	})

	// ============================================================================
	// ปุ่ม Build EXE
	// ============================================================================
	buildexe := widget.NewButton("Build EXE", func() {

		if projectPath == "" {
			logBox.SetText("❌ select folder first")
			return
		}

		//  run script
		go buildexe(projectPath, logBox)

		logBox.SetText("🚀 Build started in terminal...")
	})

	// ============================================================================
	// ปุ่มเพิ่มวัน เวลา
	// ============================================================================
	nowBtn := widget.NewButton("4 - กด เพื่อ ใส่เวลาปัจจุบัน", func() {
		now := time.Now()

		years.SetText(now.Format("2006")) //ปี
		month.SetText(now.Format("01"))   //เดือน
		days.SetText(now.Format("02"))    //วัน

		//years.SetText(now.Format("02"))   //วัน
		//timeh.SetText(now.Format("16")) //ชั่วโมง
		//timem.SetText(now.Format("04")) //นาที
	})
	// ============================================================================
	// จัดหน้ามัน
	// ============================================================================

	ui := container.NewVBox(

		container.NewVBox(
			container.NewGridWithColumns(3, selectBtn, genscripiconsBtn, buildIconsBtn),
			name, appID),
		companyName, license,

		version,
		fileversion,

		years, //month, days,
		nowBtn,

		genscripexeBtn,

		buildexe,

		logBox,
	)

	w.SetContent(ui)
	w.Resize(fyne.NewSize(500, 600))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
