package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func gallery() {
	window1 := myApp.NewWindow("Gallery")
	window1.Resize(fyne.NewSize(450, 450))

	gallery, _ := fyne.LoadResourceFromPath("./picture.png")
	window1.SetIcon(gallery)

	root_src := "C:/Users/ASUS/Desktop/DiceIcons"
	files, _ := ioutil.ReadDir(root_src)

	tabs := container.NewAppTabs()
	for _, file := range files {

		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpg" || extension == "jpeg" {
				image := canvas.NewImageFromFile(root_src + "/" + file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))

			}
		}
	}
	browsemenuitem := fyne.NewMenuItem("Browse...", func() {
		fileDialog := dialog.NewFileOpen(
			// _ to ignore error
			func(view fyne.URIReadCloser, _ error) {
				// reader to read data
				data, _ := ioutil.ReadAll(view)
				res := fyne.NewStaticResource(view.URI().Name(), data)

				img := canvas.NewImageFromResource(res)

				w := fyne.CurrentApp().NewWindow("Photos - " + view.URI().Name())
				w.SetContent(img)

				w.Resize(fyne.NewSize(1080, 520))
				w.Show() // display our image
			}, window1)

		fileDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		fileDialog.Show()
	})

	tabs.SetTabLocation(container.TabLocationLeading)

	browsemenu := fyne.NewMenu("File", browsemenuitem)
	menu := fyne.NewMainMenu(browsemenu)
	window1.SetMainMenu(menu)
	window1.SetContent(tabs)
	window1.Show()
}
