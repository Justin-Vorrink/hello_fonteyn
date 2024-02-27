package main

// Importeer IO en tijd package
import (
	"fmt"
	"time"
)

func main() {
	var kenteken string
	var kentekenList = [3]string{"15-FL-LF", "77-ll-DD", "K-456-GH"} //Maak een array aan met 3 static kentekens
	fmt.Println("Wat is uw kenteken?")
	fmt.Scanln(&kenteken) //Leest wat je typt en stopt dat in de variable "kenteken"

	//Via de functie isInKentekenList checkt hij ofdat in de kentekenList array staat
	if isInKentekenList(kenteken, kentekenList) {
		CurrentTime := time.Now()         //haalt de tijd en datum op
		CurrentHour := CurrentTime.Hour() //zet de opgehaalde tijd en datum op naar alleen het uur

		var Greeting string

		//switch case die verschillende berichten laat zien op basis van tijd.
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
	} else {
		fmt.Println("Uw kenteken staat niet in de lijst.")
	}
}

// functie om te checken ofdat het opgegeven kenteken in de array zit
func isInKentekenList(kenteken string, kentekenList [3]string) bool {
	for _, value := range kentekenList {
		if value == kenteken {
			return true
		}
	}
	return false
}
