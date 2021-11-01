package main

import (
	"encoding/json"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func toDo() {

	type Tasks struct {
		Title string
		Task  string
	}

	var myMainData []Tasks

	data_file, _ := ioutil.ReadFile("file.json")

	json.Unmarshal(data_file, &myMainData)

	// a := app.New()
	w := myApp.NewWindow("To-Do-App")
	w.Resize(fyne.NewSize(400, 400))

	icon,_ := fyne.LoadResourceFromPath("./todo.png")
	w.SetIcon(icon)

	lab_title := widget.NewLabel("***")
	lab_title.TextStyle = fyne.TextStyle{Bold: true}

	lab_task := widget.NewLabel("---")

	e_title := widget.NewEntry()
	e_title.SetPlaceHolder("Enter Your Title")

	e_task := widget.NewMultiLineEntry()
	e_task.SetPlaceHolder("Enter Your Task")

	submitBtn := widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {

		myData1 := &Tasks{
			Title: e_title.Text,
			Task:  e_task.Text,
		}

		myMainData = append(myMainData, *myData1)

		final_data, _ := json.MarshalIndent(myMainData, "", " ")

		ioutil.WriteFile("config.json", final_data, 0644)

		e_title.Text = ""
		e_task.Text = ""

		e_title.Refresh()
		e_task.Refresh()
	})

	list := widget.NewList(
		func() int { return len(myMainData) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(myMainData[lii].Title)
		},
	)

	del_btn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {

		var tempdata []Tasks

		for _, element := range myMainData {
			if lab_title.Text != element.Title {
				tempdata = append(tempdata, element)
			}
		}

		myMainData = tempdata

		result, _ := json.MarshalIndent(myMainData, "", " ")

		ioutil.WriteFile("config.json", result, 0644)

	})

	update_btn := widget.NewButtonWithIcon("", theme.UploadIcon(), func() {

		var tempdata []Tasks
		update := &Tasks{
			Title: e_title.Text,
			Task:  e_task.Text,
		}

		for _, element := range myMainData {

			if lab_title.Text == element.Title {
				tempdata = append(tempdata, *update)
			} else {
				tempdata = append(tempdata, element)
			}
		}

		myMainData = tempdata

		result, _ := json.MarshalIndent(myMainData, "", " ")

		ioutil.WriteFile("config.json", result, 0644)

		e_task.Text = ""
		e_title.Text = ""

		e_task.Refresh()
		e_title.Refresh()

		list.Refresh()
	})

	cont := container.NewHSplit(
		list,
		container.NewVBox(lab_title, lab_task, e_title, e_task,
			container.NewHBox(submitBtn, update_btn, del_btn)),
	)

	list.OnSelected = func(id widget.ListItemID) {
		lab_title.Text = myMainData[id].Title
		lab_title.Refresh()

		lab_task.Text = myMainData[id].Task
		lab_task.Refresh()
	}

	w.SetContent(cont)
	w.Show()
}
