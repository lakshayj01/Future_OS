package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var city string = "Pune"

func weather_app() {

	// a := app.New()
	w := myApp.NewWindow("Weather")
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 400))
	w.SetPadded(false)

	icon,_ := fyne.LoadResourceFromPath("./weather.png")
	w.SetIcon(icon)

	img := canvas.NewImageFromFile("./weather1.jpg")
	img.Resize(fyne.NewSize(300, 500))

	//Api part

	res, err := http.Get(fmt.Sprint("https://api.openweathermap.org/data/2.5/weather?q=", city, "&APPID=88a3325d8b543b9103c71abe0ebc15ef"))

	if err != nil {
		panic(err)
	}

	differentCities := widget.NewSelect([]string{"Pune", "Noida", "Mumbai", "Delhi", "Bangalore"}, func(str string) {
		city = str
		locations(w)
	})

	differentCities.Resize(fyne.NewSize(200, 10))
	differentCities.PlaceHolder = city

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		panic(err)
	}

	temperature := canvas.NewText(fmt.Sprint(int32((weather.Main.Temp) - 273.15)),color.Black)
	temperature.TextSize = 80
	temperature.TextStyle = fyne.TextStyle{Bold : true,Italic: true}

	city_label := canvas.NewText(fmt.Sprint(weather.Sys.Country,", ",weather.Name,", °C"),color.Black)
	log.Println(weather.Wind.Speed)
	wind_label := canvas.NewText(fmt.Sprint("Wind Speed ",int64(weather.Wind.Deg),"°"),color.Black)
	
	w.SetContent(container.NewBorder(nil,nil,nil,nil,img,container.NewVBox(	differentCities,container.NewCenter(temperature),
	container.NewCenter(city_label),
	container.NewCenter(wind_label),

	)))



	w.Show()
}

func locations(w fyne.Window) {
	img := canvas.NewImageFromFile("./weather1.jpg")
	switch city {
	case "Delhi":
		img = canvas.NewImageFromFile("./weather2.jpg")
	case "Noida":
		img = canvas.NewImageFromFile("./weather3.jpg")
	case "Bangalore":
		img = canvas.NewImageFromFile("./weather4.jpg")
	case "Mumbai":
		img = canvas.NewImageFromFile("./weather5.jpg")

	}
	img.Resize(fyne.NewSize(300, 500))

	res, err := http.Get(fmt.Sprint("https://api.openweathermap.org/data/2.5/weather?q=", city, "&APPID=88a3325d8b543b9103c71abe0ebc15ef"))

	if err != nil {
		panic(err)
	}

	differentCities := widget.NewSelect([]string{"Pune", "Noida", "Mumbai", "Delhi", "Bangalore"}, func(str string) {
		city = str
		locations(w)
	})

	differentCities.Resize(fyne.NewSize(200, 10))
	differentCities.PlaceHolder = city

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		panic(err)
	}

	temperature := canvas.NewText(fmt.Sprint(int32((weather.Main.Temp) - 273.15)),color.Black)
	temperature.TextSize = 90
	temperature.TextStyle = fyne.TextStyle{Bold : true,Italic: true}

	city_label := canvas.NewText(fmt.Sprint(weather.Sys.Country,", ",weather.Name,", °C"),color.Black)
	log.Println(weather.Wind.Speed)
	wind_label := canvas.NewText(fmt.Sprint("Wind Speed ",int64(weather.Wind.Deg),"°"),color.Black)


	w.SetContent(container.NewBorder(nil,nil,nil,nil,img,container.NewVBox(differentCities,container.NewCenter(temperature),
	container.NewCenter(city_label),
	container.NewCenter(wind_label),

	
	)))
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
