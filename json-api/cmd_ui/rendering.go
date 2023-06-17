package rendering

import (
	"fmt"
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
