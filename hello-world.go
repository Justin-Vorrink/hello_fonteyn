package main

// Importeer IO en tijd package
import (
	"fmt"
	"time"
)

func main() {

	CurrentTime := time.Now()
	CurrentHour := CurrentTime.Hour()

	var Greeting string

	switch {
	case CurrentHour >= 7 && CurrentHour < 12:
		Greeting = "Goedemorgen! Welkom bij Fonteyn Vakantieparken"
	case CurrentHour >= 12 && CurrentHour < 18:
		Greeting = "Goedemiddag! Welkom bij Fonteyn Vakantieparken"
	case CurrentHour >= 18 && CurrentHour < 23:
		Greeting = "Goedenavond! Welkom bij Fonteyn Vakantieparken"
	default:
		Greeting = "Sorry, de parkeerplaats is ’s nachts gesloten!"
	}
	fmt.Println(Greeting)

}
