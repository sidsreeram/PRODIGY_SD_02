package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1
	attempts := 0

	
	a := app.New()
	w := a.NewWindow("Guess the Number")

	
	instruction := widget.NewLabel("I have picked a random number between 1 and 100. Can you guess it?")
	feedback := widget.NewLabel("Make your guess:")
	guessEntry := widget.NewEntry() 
	guessEntry.SetPlaceHolder("Enter your guess here")
	attemptLabel := widget.NewLabel("Attempts: 0")

	
	var submitButton *widget.Button
	submitButton = widget.NewButton("Submit", func() {
		attempts++

		guess, err := strconv.Atoi(guessEntry.Text)
		if err != nil {
			feedback.SetText("Please enter a valid number!")
			return
		}

		if guess < randomNumber {
			feedback.SetText("Too low! Try again.")
		} else if guess > randomNumber {
			feedback.SetText("Too high! Try again.")
		} else {
			feedback.SetText(fmt.Sprintf("Congrats! You guessed the number in %d attempts.", attempts))
			submitButton.Disable() 
		}

	
		attemptLabel.SetText(fmt.Sprintf("Attempts: %d", attempts))
	})

	content := container.NewVBox(instruction, guessEntry, submitButton, feedback, attemptLabel)


	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 200)) 
	w.ShowAndRun()
}
