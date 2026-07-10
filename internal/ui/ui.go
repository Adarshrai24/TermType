package ui 

import(
	"fmt"
	"github.com/Adarshrai24/ttyper/internal/storage"
)

func ShowMainMenu() int {
	fmt.Println("1. Start Test")
	fmt.Println("2. History")
	fmt.Println("3. Statistics")
	fmt.Println("4. Exit")
	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scanf("%d", &choice)
	return choice
}

func ShowTimeMenu() int {
	fmt.Println("1. 15sec")
	fmt.Println("2. 30sec")
	fmt.Println("3. 60sec")
	fmt.Println("4. 120sec")
	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		return 15
	case 2:
		return 30
	case 3:
		return 60
	case 4:
		return 120
	default:
		return 30
	}
}

func PrintHistory(history []storage.Session) {
	fmt.Printf("%-10s %-10s %-12s %-10s\n", "GWPM", "NWPM", "Accuracy", "Duration")

	fmt.Println("------------------------------------------------")

	for _, session := range history {
		fmt.Printf("%-10.2f %-10.2f %-12.2f %-10d\n",
			session.GWPM,
			session.NWPM,
			session.Accuracy,
			session.Duration,
		)
	}
}

func PrintMaxStats() {
	fmt.Println("----------------MaximumStats--------------------")
}

func PrintAvgStats() {
	fmt.Println("----------------AverageStats--------------------")
}

func PrintStats(stats []storage.Stats) {
	fmt.Printf("%-10s %-10s %-10s %-10s\n",
		"Duration", "GWPM", "NWPM", "Accuracy")
	fmt.Println("------------------------------------------------")

	for _, s := range stats {
		fmt.Printf("%-10d %-10.2f %-10.2f %-10.2f\n",
			s.Duration,
			s.GWPM,
			s.NWPM,
			s.Accuracy,
		)
	}
}


