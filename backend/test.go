package main

import (
	"fmt"
	"time"
)

func main(){

	var downloadState int
	downloadState = 1

	for ok := true; ok; ok = downloadState != 2 {
		//n, err := downloadState
		if downloadState < 1  {
			fmt.Println("invalid input")
			break
		}
		time.Sleep(2 * time.Second)
		switch downloadState {
		case 1:
			fmt.Println("cause 1")
			downloadState = 5
		case 5:
			fmt.Println("case 5")
			// Do nothing (we want to exit the loop)
			// In a real program this could be cleanup
		default:
			fmt.Println("not clear")
		}
	}

}
