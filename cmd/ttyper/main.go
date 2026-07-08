package main 

import(
	"fmt"
	"github.com/Adarshrai24/ttyper/internal/ui"
	"github.com/Adarshrai24/ttyper/internal/data"
	"github.com/Adarshrai24/ttyper/internal/test"
)

func main() {
	for {
		fmt.Println("---------------Welcome to ttyper------------------")
		
		choice := ui.ShowMainMenu()

		if choice == 1 {
			timeChoice := ui.ShowTimeMenu()
			passageset := data.RandomParagraphPick()
			result := test.Start(passageset.Passages, passageset.Current, timeChoice)
			fmt.Println("WPM: ", result.WPM)
			fmt.Println("Accuracy: ", result.Accuracy)
			fmt.Println("TimeFormat: ", timeChoice)
		} else if choice == 2 {

		} else if choice == 3 {

		} else {
			break
		}
	}
}
