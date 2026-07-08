package test 

import (
	"fmt"
)

type Result struct {
	WPM int 
	Accuracy float64
}

func Start(passages []string, index int, duration int) Result {
	fmt.Println(passages[index])		

	result := Result{
		WPM: 0,
		Accuracy: 0,
	}

	return result
}
