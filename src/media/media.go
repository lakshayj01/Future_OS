package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	a := app.New()
	window4 := a.NewWindow("Media Player")
	window4.Resize(fyne.NewSize(400, 400))

	icon,_ := fyne.LoadResourceFromPath("./spotify.png")
	window4.SetIcon(icon)

	logo := canvas.NewImageFromFile("./spotify.png")
	logo.FillMode = canvas.ImageFillOriginal

	filepath := widget.NewEntry()

	lblTimeUsed := widget.NewLabel("")
	lblTimeUsed.Alignment = fyne.TextAlignCenter

	progress := widget.NewProgressBar()
	progress.Min = 0
	progress.Max = 100
	progress.Value = 0

	done := make(chan bool)
	var ctrl *beep.Ctrl
	var volume *effects.Volume

	bpath := binding.NewString()
	bpath.AddListener(binding.NewDataListener(func() {
		ctrl = nil
	}))
	bpath.Set("C:\\Users\\ASUS\\Desktop\\DiceImg\\Sample1.mp3")

	openFile := widget.NewButtonWithIcon("", theme.SearchIcon(), func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {

			filepath.Text = uc.URI().Name()
			bpath.Set(uc.URI().Path())
			filepath.Refresh()
		}, window4)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})

	var playbtn *widget.Button

	playbtn = widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		if ctrl != nil {
			if !ctrl.Paused {
				playbtn.SetIcon(theme.MediaPlayIcon())
			} else {
				playbtn.SetIcon(theme.MediaPauseIcon())
			}
			speaker.Lock()
			ctrl.Paused = !ctrl.Paused
			speaker.Unlock()

		} else {

			playbtn.SetIcon(theme.MediaPauseIcon())
			pth, _ := bpath.Get()
			go func() {

				f, err := os.Open(pth)
				if err != nil {
					panic(err)
				}

				streamer, format, err := mp3.Decode(f)
				if err != nil {
					panic(err)
				}

				defer streamer.Close()

				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				ctrl = &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
				progress.Max = float64(streamer.Len())

				volume = &effects.Volume{
					Streamer: ctrl,
					Base:     2,
					Volume:   0,
					Silent:   false,
				}
				speaker.Play(volume)

				for {
					select {
					case <-done:
						speaker.Clear()
						return
					case <-time.After(time.Second):
						progress.SetValue(float64(streamer.Position()))
						speaker.Lock()
						pos := format.SampleRate.D(streamer.Position())
						lenn := format.SampleRate.D(streamer.Len())
						speaker.Unlock()
						lblTimeUsed.SetText(fmt.Sprintf("%v / %v", pos.Round(time.Second), lenn.Round(time.Second)))
					}
				}

			}()
		}
	})

	stopbtn := widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() {
		done <- false
		ctrl = nil
		lblTimeUsed.SetText("")
		progress.SetValue(0)
		playbtn.SetIcon(theme.MediaPlayIcon())

	})

	volSlider := widget.NewSlider(-10, 10)
	volSlider.SetValue(10)
	volSlider.Orientation = widget.Vertical

	volSlider.OnChanged = func(f float64) {

		if volume == nil {
			return
		}
		speaker.Lock()
		volume.Volume = f
		speaker.Unlock()
	}
	
	// var volbtn *widget.Button
	// volbtn = widget.NewButtonWithIcon("", theme.VolumeUpIcon(), func() {
	// 	widget.ShowPopUpAtPosition(volSlider, fyne.CurrentApp().Driver().CanvasForObject(volbtn), volbtn.Position().Add(fyne.NewDelta(volbtn.Size().Width/2, 0)))
	// })
	var volbtn *contextMenuButton
	volbtn = newContextmenuButton(theme.VolumeUpIcon(),volSlider)
	
	hbox := container.NewHBox(layout.NewSpacer(),openFile, playbtn, stopbtn, volbtn,layout.NewSpacer())
	top := container.NewVBox(logo,filepath, progress,lblTimeUsed, hbox)
	cont := container.NewBorder(top, nil, nil, nil)
	window4.SetContent(cont)
	window4.CenterOnScreen()
	window4.ShowAndRun()
}

type contextMenuButton struct{
	widget.Button
	menu fyne.CanvasObject
}

func (b *contextMenuButton) Tapped(e *fyne.PointEvent) {
	widget.ShowPopUpAtPosition(b.menu,fyne.CurrentApp().Driver().CanvasForObject(b),e.AbsolutePosition.Subtract(fyne.NewDelta(0,b.menu.Size().Height)))	
}

func newContextmenuButton(icon fyne.Resource,menu fyne.CanvasObject) *contextMenuButton  {
	b := &contextMenuButton{menu : menu}
	b.Icon = icon

	b.ExtendBaseWidget(b)
	return b
}