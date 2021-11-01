package main

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWin fyne.Window = myApp.NewWindow("Welcome")

var calcbtn fyne.Widget
var txt_editorbtn fyne.Widget
var txt_to_Pdfbtn fyne.Widget
var weatherbtn fyne.Widget
var stockbtn fyne.Widget
var themeBtn fyne.Widget
var mp3btn fyne.Widget
var picbtn fyne.Widget

func main() {

	myWin.Resize(fyne.NewSize(250, 250))

	label := widget.NewLabel("Enter Your credentials")

	mainicon, _ := fyne.LoadResourceFromPath("./peplogo.png")
	myWin.SetIcon(mainicon)

	entry1 := widget.NewEntry()
	password1 := widget.NewPasswordEntry()

	form := widget.NewForm(

		widget.NewFormItem("Username", entry1),

		widget.NewFormItem("Password", password1),
	)

	form.OnCancel = func() {
		label.Text = "Canceled"
		label.Refresh()
	}

	form.OnSubmit = func() {

		if entry1.Text == "qwerty" && password1.Text == "1234" {
			label.Text = "Welcome to os"
			ap()
			myWin.Close()
		} else {
			label.Text = "Invalid credentials"
			label.Refresh()

		}
	}

	myWin.SetContent(
		container.NewVBox(
			label,
			form,
		),
	)
	myWin.ShowAndRun()
}

func ap() {

	mainw := fyne.CurrentApp().NewWindow("Future OS")
	mainw.Resize(fyne.NewSize(1280, 700))
	istheme := true

	deskicon, _ := fyne.LoadResourceFromPath("./desktop.png")
	mainw.SetIcon(deskicon)

	img := canvas.NewImageFromURI(storage.NewFileURI("./bg.jpg"))
	img.ScaleMode = canvas.ImageScale(canvas.ImageFillOriginal)

	calicon, _ := fyne.LoadResourceFromPath("./calculator.png")
	calcbtn = widget.NewButtonWithIcon("", calicon, func() {
		calculator()
	})

	txticon, _ := fyne.LoadResourceFromPath("./files.png")
	txt_editorbtn = widget.NewButtonWithIcon("", txticon, func() {
		text_app()
	})

	pdficon, _ := fyne.LoadResourceFromPath("./pdf.png")
	txt_to_Pdfbtn = widget.NewButtonWithIcon("", pdficon, func() {
		txt_convertor()
	})

	mp3icon, _ := fyne.LoadResourceFromPath("./spotify.png")
	mp3btn = widget.NewButtonWithIcon("", mp3icon, func() {
		multimedia()
	})

	picicon, _ := fyne.LoadResourceFromPath("./picture.png")
	picbtn = widget.NewButtonWithIcon("", picicon, func() {
		gallery()
	})

	wthricon, _ := fyne.LoadResourceFromPath("./weather.png")
	weatherbtn = widget.NewButtonWithIcon("", wthricon, func() {
		weather_app()
	})

	todoicon, _ := fyne.LoadResourceFromPath("./todo.png")
	todobtn := widget.NewButtonWithIcon("",todoicon,func() {
		toDo()
	})

	themeicon, _ := fyne.LoadResourceFromPath("./themes.png")
	themebtn := widget.NewButtonWithIcon("", themeicon, func() {
		if istheme {
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
			istheme = false
		} else {
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
			istheme = true
		}
	})
	istheme = !istheme

	quiticon, _ := fyne.LoadResourceFromPath("./home.png")
	quitbtn := widget.NewButtonWithIcon("", quiticon, func() {
		myApp.Quit()
	})

	url, _ := url.Parse("https://www.pepcoding.com/")
	link := widget.NewHyperlinkWithStyle("Github", url, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	clock := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{
		Bold:      true,
		Italic:    true,
		Monospace: false,
		TabWidth:  2,
	})

	showTime(clock)
	go func() {
		tick := time.NewTicker(time.Second)

		for range tick.C {
			showTime(clock)
		}
	}()

	hbox := container.NewHBox(calcbtn, picbtn, txt_editorbtn, weatherbtn, txt_to_Pdfbtn, mp3btn,todobtn, layout.NewSpacer(), quitbtn, themebtn, link, clock)
	top := container.NewVBox(hbox)
	cont := container.NewBorder(nil, top, nil, nil, img)

	mainw.FullScreen()
	mainw.SetPadded(false)
	mainw.CenterOnScreen()

	mainw.SetContent(cont)
	mainw.Show()
}

func showTime(clock *widget.Label) {

	formatted := time.Now().Format("02/01 Monday 15:04:05")
	clock.SetText(formatted)
}
