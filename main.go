package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Gopher")

	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { myApp.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Wilbur"),
			), myWindow)
		}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	myWindow.SetMainMenu(mainMenu)

	// Define a welcome text centered
	text := canvas.NewText("Display a random Gopher!", color.White)
	text.Alignment = fyne.TextAlignCenter

	// Define a Gopher image
	var resource, _ = fyne.LoadResourceFromURLString("https://dkrn4sk0rn31v.cloudfront.net/uploads/2022/10/o-que-e-e-como-comecar-com-golang.jpg")
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500}) // by default size is 0, 0

	// Define a "random" button
	images := []string{
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS2toMici6yMVKrZ0DpeSFWpagSHdiK_aQ63iPYj1KaJDA__S_XQgigPMIwLYqOENdXTf8&usqp=CAU",
		"https://fingers-site-production.s3.eu-central-1.amazonaws.com/uploads/images/NYgf3QgEdUpHTR7YIYacJanBU3JEeDxmIKGOUKcD.jpg",
		"https://dce0qyjkutl4h.cloudfront.net/wp-content/uploads/2020/10/golang-Programing.jpg",
	}
	randomBtn := widget.NewButton("Random", func() {
		randomNumber := rand.Intn(3)
		fmt.Println(randomNumber)

		resource, _ := fyne.LoadResourceFromURLString(images[randomNumber])
		gopherImg.Resource = resource

		//Redrawn the image with the new path
		gopherImg.Refresh()
	})
	randomBtn.Importance = widget.HighImportance
	randomBtn.Alignment = widget.ButtonAlignCenter

	// Display a vertical box containing text, image and button
	box := container.NewVBox(
		text,
		gopherImg,
		randomBtn,
	)

	// Display our content
	myWindow.SetContent(box)

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	// Show window and run app
	myWindow.ShowAndRun()
}
