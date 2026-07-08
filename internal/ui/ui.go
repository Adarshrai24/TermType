package ui 

import(
	"fmt"
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
