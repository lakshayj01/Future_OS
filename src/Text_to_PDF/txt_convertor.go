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
	"github.com/jung-kurt/gofpdf"
)

func main() {

	a := app.New()
	w := a.NewWindow("TextFile to PDF")
	w.Resize(fyne.NewSize(450, 450))

	pdf_opener, _ := fyne.LoadResourceFromPath("./pdf.png")
	w.SetIcon(pdf_opener)
	label3 := widget.NewLabel("Thanks for Watching :)")

	flag := 1

	file_Dialog := dialog.NewFileOpen(
		func(read fyne.URIReadCloser, _ error) {
			data, _ := ioutil.ReadAll(read)

			result := fyne.NewStaticResource(read.URI().Name(), data)

			entry := widget.NewMultiLineEntry()

			entry.SetText(string(result.StaticContent))

			w := fyne.CurrentApp().NewWindow(string(result.StaticName) + " - Notepad")

			w.SetContent(container.NewScroll(entry))
			w.Resize(fyne.NewSize(400, 400))

			menuItem1 := fyne.NewMenuItem("Create pdf", func() {
				fsdialog := dialog.NewFileSave(
					func(uc fyne.URIWriteCloser, e error) {
						wrtdata := string(entry.Text)

						pdf := gofpdf.New("P", "mm", "A4", "")
						pdf.AddPage()
						pdf.SetFont("Arial", "B", 16)
						pdf.MultiCell(190, 5, wrtdata, "2", "2", false)
						_ = pdf.OutputFileAndClose("F:\\PDF's\\sample" + strconv.Itoa(flag) + ".pdf")
						flag++

						uc.Write([]byte(wrtdata))
					}, w)
				fsdialog.SetFileName("N_File" + strconv.Itoa(flag) + ".txt")
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

	w.SetContent(label3)

	w.ShowAndRun()
}