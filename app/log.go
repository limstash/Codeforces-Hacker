package app

import (
	"fmt"
	"os"
	"time"
)

func log(level int, text string) {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05"))

	if level == 1 {
		fmt.Printf(" [Error] " + text + "\n")

		time.Sleep(time.Millisecond * 3000)
		os.Exit(1)
	} else if level == 2 {
		fmt.Printf(" [Warning] " + text + "\n")
	} else {
		fmt.Printf(" [Notice] " + text + "\n")
	}
}
