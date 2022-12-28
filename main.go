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

//sleep for a 25 mins
func pomidorWork() {
	time.Sleep(25 * time.Minute)
}

//create Work window
func showWorkWindow() {
	a := app.New()

	w := a.NewWindow("Time to work")

	msg := widget.NewLabel("25 minustes work")

	w.SetContent(msg)
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
	time.Sleep(2 * time.Second)
	w.Hide()
}

//create Rest Window
func showRestWindow() {
	a := app.New()

	w := a.NewWindow("Rest")

	msg := widget.NewLabel("You have 5 minustes rest")

	w.SetContent(msg)

	w.Resize(fyne.NewSize(100, 100))

	w.ShowAndRun()

	w.Close()

}

//sleep 5 mins - rest
func pomidorRest() {

	time.Sleep(5 * time.Minute)

}

func main() {

	var hours int = 8

	fmt.Print("Input the number of hours you want to study: ... ")

	numberOfPomidors := countPomidors(hours)

	fmt.Print(numberOfPomidors)

	for i := 0; i < numberOfPomidors; i++ {
		showWorkWindow()

		pomidorWork()

		showRestWindow()

		pomidorRest()

	}
}
