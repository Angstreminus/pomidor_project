package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func countPomidors(hours int) int {
	return (hours * 60) / 30
}

func pomidorWork() {
	a := app.New()

	w := a.NewWindow("Time to work")

	msg := widget.NewLabel("25 minustes work")

	w.SetContent(msg)
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
	time.Sleep(2 * time.Second)
	w.Hide()
	time.Sleep(25 * time.Minute)

}

func pomidorRest() {

	a := app.New()

	w := a.NewWindow("Rest")

	msg := widget.NewLabel("You have 5 minustes rest")

	w.SetContent(msg)
	w.Resize(fyne.NewSize(100, 100))

	w.ShowAndRun()
	w.Close()
	time.Sleep(3 * time.Second)

	time.Sleep(5 * time.Minute)

}

func main() {

	var (
		hours int
	)

	fmt.Print("Input the number of hours you want to study: ... ")

	hours = 8

	numberOfPomidors := countPomidors(hours)

	fmt.Print(numberOfPomidors)

	for i := 0; i < numberOfPomidors; i++ {
		pomidorWork()

		pomidorRest()

	}
}
