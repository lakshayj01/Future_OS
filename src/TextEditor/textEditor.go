package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)
var count = 1
func main() {
	
	a := app.New()
	w := a.NewWindow("Welcome to Text_Editor")
	w.Resize(fyne.NewSize(450, 450))

	text_app, _ := fyne.LoadResourceFromPath("./files.png")
	w.SetIcon(text_app)

	cont := container.NewVBox()

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Wrapping = fyne.TextWrapBreak

	btm1 := widget.NewButton("Add More file", func() {
		cont.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	})

	btm3 := widget.NewButton("Save", func() {
		fileSaveDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)
				uc.Write(textData)
			}, w)
		fileSaveDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")

		fileSaveDialog.Show()
	})

	btm2 := widget.NewButton("Open Text files", func() {
		file_Dialog := dialog.NewFileOpen(
			func(read fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(read)

				result := fyne.NewStaticResource(read.URI().Name(), data)

				entry := widget.NewMultiLineEntry()

				entry.SetText(string(result.StaticContent))

				w := fyne.CurrentApp().NewWindow(string(result.StaticName) + " - Notepad")

				w.SetContent(container.NewScroll(entry))
				w.Resize(fyne.NewSize(400, 400))

				menuItem1 := fyne.NewMenuItem("Save", func() {
					fsdialog := dialog.NewFileSave(
						func(uc fyne.URIWriteCloser, _ error) {
							wrtdata := []byte(entry.Text)
							uc.Write(wrtdata)
						}, w)
					fsdialog.SetFileName("NewFile.txt")
					fsdialog.Show()
				})

				newMenu := fyne.NewMenu("File", menuItem1)
				menu := fyne.NewMainMenu(newMenu)

				w.SetMainMenu(menu)

				w.Show()
			}, w)

		file_Dialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))
		file_Dialog.Show()

	})

	cont.Add(btm1)

	content := container.NewHSplit(
		container.NewVBox(
			cont,
			container.NewGridWithRows(2,

				btm2,
				btm3,
			),
		),
		container.NewVBox(
			input,
		),
	)

	w.SetContent(content)
	w.ShowAndRun()
}