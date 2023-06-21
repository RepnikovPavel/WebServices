package cmd_ui

import (
	"fmt"
	"log"
	"time"
)

func Spinner() {
	fmt.Printf("\n")
	FrameDuration := 67 * time.Millisecond
	for {
		for _, char := range `\|/-` {
			fmt.Printf("\r%c", char)
			time.Sleep(FrameDuration)
		}
	}
}

func LOGERR(err error) {
	log.Printf("error: %s\n", err)
}
