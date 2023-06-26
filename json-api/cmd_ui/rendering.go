package cmd_ui

import (
	"errors"
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

func LOGERRSEQ(errs ...error) {
	log.Printf("\nerror seq: \n%s\n", errors.Join(errs...))
}
