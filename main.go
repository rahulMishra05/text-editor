package main

import (
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var count int = 1


func main() {
	a := app.New()

	w := a.NewWindow("Text Editor")
	w.Resize(fyne.NewSize(600,600))


	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Text Editor"),
		),
	)


	content.Add(widget.NewButton("Add New File", func()  {
		content.Add(widget.NewLabel("New File"+ strconv.Itoa(count)))
		count++
	}))


	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text......")
	input.Resize(fyne.NewSize(400,400))



	saveBtn := widget.NewButton("Save text file", func() {

	})


	w.SetContent(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
				saveBtn,
			),
		),
	)

	w.ShowAndRun()
}