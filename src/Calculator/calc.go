package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {

	a := app.New()
	calw := a.NewWindow("Calculator")

	istheme := true
	calculator, _ := fyne.LoadResourceFromPath("./calculator.png")
	calw.SetIcon(calculator)

	output := ""
	input := widget.NewLabelWithStyle(output, fyne.TextAlignLeading, fyne.TextStyle{
		Bold:      true,
		Italic:    true,
		Monospace: false,
		TabWidth:  5,
	})

	hisString := ""
	history := widget.NewLabel(hisString)
	var hisArr = []string{}
	isHistory := false

	brown_btn := canvas.NewRectangle(
	color.NRGBA{R: 101, G: 70, B: 33, A: 255})

	clearBTn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		output = ""
		input.SetText(output)
	})
	contx := container.New(
		layout.NewMaxLayout(),
		brown_btn,
		clearBTn,
	)

	backBTn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		if len(output) > 0 {

			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	conty := container.New(
		layout.NewMaxLayout(),
		brown_btn,
		backBTn,
	)

	orange_btn := canvas.NewRectangle(
		color.NRGBA{R: 255, G: 170, B: 0, A: 255})

	divBTn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})
	cont1 := container.New(
		layout.NewMaxLayout(),
		orange_btn,
		divBTn,
	)

	multiBTn := widget.NewButtonWithIcon("",theme.CancelIcon() ,func() {
		output = output + "*"
		input.SetText(output)
	})
	cont2 := container.New(
		layout.NewMaxLayout(),
		orange_btn,
		multiBTn,
	)

	minusBTn := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(),func() {
		output = output + "-"
		input.SetText(output)
		
	})
	cont3 := container.New(
		layout.NewMaxLayout(),
		orange_btn,
		minusBTn,
	)

	plusBTn := widget.NewButtonWithIcon("", theme.ContentAddIcon(),func() {
		output = output + "+"
		input.SetText(output)
	})
	cont4 := container.New(
		layout.NewMaxLayout(),
		orange_btn,
		plusBTn,
	)

	sevenBTn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	eightBTn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	nineBTn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	fourBTn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	fiveBTn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	sixBTn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	oneBTn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	twoBTn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	threeBTn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	zeroBTn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	deciBTn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	openBTn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	closeBTn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})

	isPrimeBTn := widget.NewButton("isPrime", func() {

		if len(output) > 0 {

			ans, err := strconv.Atoi(output)
			if err != nil {
				panic(err)
			}
			if isPrime(ans) == 1 {
				input.SetText("TRUE")
			} else {
				input.SetText("FALSE")
			}
		}
	})

	historyBTn := widget.NewButtonWithIcon("", theme.HistoryIcon(), func() {

		if isHistory {
			hisString = ""
		} else {
			for i := len(hisArr) - 1; i >= 0; i-- {
				hisString = hisString + hisArr[i]
				hisString += "\n"
			}
		}
		isHistory = !isHistory

		history.SetText(hisString)
	})
	contz := container.New(
		layout.NewMaxLayout(),
		brown_btn,
		historyBTn,
	)

	blue_btn := canvas.NewRectangle(
		color.NRGBA{R: 0, G: 180, B: 255, A: 255})

	eqaulBTn := widget.NewButton("=", func() {

		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				elementToAppend := output + "=" + ans
				hisArr = append(hisArr, elementToAppend)
				output = ans
			} else {
				output = "ERROR: Press the trash button"
			}
		} else {
			output = "ERROR: Press the trash button"
		}
		input.SetText(output)

	})
	cont5 := container.New(
		layout.NewMaxLayout(),
		blue_btn,
		eqaulBTn,
	)

	
	themeBtn := widget.NewButtonWithIcon("", theme.ColorChromaticIcon(), func() {
		if istheme {
			a.Settings().SetTheme(theme.LightTheme())
			istheme = false
		} else {
			a.Settings().SetTheme(theme.DarkTheme())
			istheme = true
		}
	})
	istheme = !istheme

	btnSquare := widget.NewButton("x²",func() {
		expression, _ := govaluate.NewEvaluableExpression(fmt.Sprint(input.Text));
		parameters := make(map[string]interface{})
		result, _ := expression.Evaluate(parameters);
		newValue,_ :=   strconv.Atoi(fmt.Sprint(result))
		input.Text =fmt.Sprint(newValue*newValue)
		input.Refresh()
	})

	calw.SetContent(

		container.NewVBox(
			input,
			history,
			container.NewGridWithColumns(4,
				contx,
				conty,
				contz,
				cont1),
			container.NewGridWithColumns(4,
				sevenBTn,
				eightBTn,
				nineBTn,
				cont2),
			container.NewGridWithColumns(4,
				fourBTn,
				fiveBTn,
				sixBTn,
				cont3),
			container.NewGridWithColumns(4,
				oneBTn,
				twoBTn,
				threeBTn,
				cont4),
			container.NewGridWithColumns(4,
				isPrimeBTn,
				zeroBTn,
				deciBTn,
				cont5),
			container.NewGridWithColumns(4,
				openBTn,
				closeBTn,
				btnSquare,
				themeBtn,
				),
		),
	)
	// a.Settings().Scale()
	calw.ShowAndRun()
}

func isPrime(x int) int {

	if x == 0 || x == 1 {
		return 0
	}

	factor := 0
	for div := 2; div*div <= x; div++ {
		if x%div == 0 {
			factor++
			break
		}
	}
	if factor == 0 {
		return 1
	} else {
		return 0
	}

}
