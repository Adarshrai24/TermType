package main 

import(
	"fmt"
	"log"
	"github.com/Adarshrai24/ttyper/internal/ui"
	"github.com/Adarshrai24/ttyper/internal/data"
	"github.com/Adarshrai24/ttyper/internal/test"
	"github.com/Adarshrai24/ttyper/internal/storage"
)

func main() {
	for {
		fmt.Println("---------------Welcome to ttyper------------------")
		
		db, err := storage.Open()
		if err != nil {
			log.Fatal(err) 
		}
		defer db.Close()

		choice := ui.ShowMainMenu()
		
		if choice == 1 {
			timeChoice := ui.ShowTimeMenu()
			passageset := data.RandomParagraphPick()
			result := test.Start(passageset.Passages, passageset.Current, timeChoice)
			fmt.Println("GWPM: ", result.GWPM)
			fmt.Println("NWPM: ", result.NWPM)
			fmt.Println("Accuracy: ", result.Accuracy)
			fmt.Println("TimeFormat: ", timeChoice)
			err := storage.SaveSession(db, result, timeChoice)
			if err != nil {
				log.Fatal(err)
			}
		} else if choice == 2 {
			history, err := storage.GetHistory(db)
			if err != nil {
				log.Fatal(err)
			}
			ui.PrintHistory(history)
		} else if choice == 3 {
			// for starting I am for now ignoring errors will handle later properly
			maxStats15Sec,_ := storage.MaximumStats(db, 15)
			maxStats30Sec,_ := storage.MaximumStats(db, 30)
			maxStats60Sec,_ := storage.MaximumStats(db, 60)
			maxStats120Sec,_ := storage.MaximumStats(db, 120)
			avgStats15Sec,_ := storage.AverageStats(db, 15)
			avgStats30Sec,_ := storage.AverageStats(db, 30)
			avgStats60Sec,_ := storage.AverageStats(db, 60)
			avgStats120Sec,_ := storage.AverageStats(db, 120)
			maxStats := []storage.Stats{
				maxStats15Sec,
				maxStats30Sec,
				maxStats60Sec,
				maxStats120Sec,
			}
			avgStats := []storage.Stats{
				avgStats15Sec,
				avgStats30Sec,
				avgStats60Sec,
				avgStats120Sec,
			}
			ui.PrintMaxStats()
			ui.PrintStats(maxStats)
			ui.PrintAvgStats()
			ui.PrintStats(avgStats)
		} else {
			break
		}
	}
}
