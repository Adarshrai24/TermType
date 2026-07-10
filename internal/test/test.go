package test 

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"golang.org/x/term" // for reading character from terminal one by one 
)

/*
GWPM = (total chars)/time in secs * 12
NWPM = (total chars - mistakes)/time in secs * 12
Accuracy = (total chars - mistakes)/total chars * 100
*/

type Result struct {
	GWPM float64
	NWPM float64
	Accuracy float64
}

func Start(passages []string, index int, duration int) Result {
	fmt.Println(passages[index])		
	
	charCount := 0
	mistakes := 0
	timerStarted := false
	currentIndex := index	
	totalCharCount := 0
	//save the original state and switch to raw mode
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState) // restore the original state of terminal when program exits
	
	reader := bufio.NewReader(os.Stdin)
	
	var timer *time.Timer
	//start := time.Now()
	//typing loop
	Loop:
	for {
		r, _, _ := reader.ReadRune()
		// Ignore Backspace
		if r == '\b' || r == 127 {
			continue
		}

		// Ignore Escape sequences (arrow keys)
		if r == 27 {
			reader.ReadRune()
			reader.ReadRune()
			continue
		}

		// Ctrl+C
		if r == 3 {
			break Loop
		}	

		if !timerStarted {
			//start = time.Now()
			timer = time.NewTimer(time.Duration(duration) * time.Second)
			timerStarted = true
		}

		fmt.Printf("%c", r)
		if r != rune(passages[currentIndex][charCount]) {
			mistakes++
			continue
		}
		charCount++
		totalCharCount++
		if charCount == len(passages[currentIndex]) {
			currentIndex = (currentIndex + 1) % len(passages)
			fmt.Println(passages[currentIndex])
			charCount = 0
		}

		if timerStarted {
			select {
			case <-timer.C:
				break Loop
			default:
			}
		}	
	}
	fmt.Println()
	//fmt.Println("Elapsed:", time.Since(start))
	grossWordPerMinute := float64(totalCharCount)/float64(duration) * 12
	netWordPerMinute := float64(totalCharCount - mistakes)/float64(duration) * 12
	accuracy := float64(totalCharCount - mistakes)/float64(totalCharCount) * 100

	result := Result{
		GWPM: grossWordPerMinute,
		NWPM: netWordPerMinute,
		Accuracy: accuracy,
	}

	return result
}
