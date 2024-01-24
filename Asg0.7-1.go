package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var conversation string
	fmt.Println("Enter input string:")
	fmt.Scanln(&conversation)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		separateMessages(conversation)
	}()
	wg.Wait()
}

func separateMessages(conversation string) {
	aliceCh := make(chan string)
	bobCh := make(chan string)
	done := make(chan struct{})

	go func() {
		defer close(done)

		currentSpeaker := "alice"
		messageBuffer := ""

		for _, char := range conversation {
			switch char {
			case '$':
				if messageBuffer != "" {
					aliceCh <- messageBuffer
					messageBuffer = ""
				}
				currentSpeaker = "alice"
			case '#':
				if messageBuffer != "" {
					bobCh <- messageBuffer
					messageBuffer = ""
				}
				currentSpeaker = "bob"
			case '^':
				return
			default:
				messageBuffer += string(char)
			}
		}

		if messageBuffer != "" {
			switch currentSpeaker {
			case "alice":
				aliceCh <- messageBuffer
			case "bob":
				bobCh <- messageBuffer
			}
		}

		close(aliceCh)
		close(bobCh)
	}()

	for {
		select {
		case msg, ok := <-aliceCh:
			if !ok {
				aliceCh = nil
			} else {
				fmt.Print("alice : ", msg, ",")
			}
		case msg, ok := <-bobCh:
			if !ok {
				bobCh = nil
			} else {
				fmt.Print("bob : ", msg, ",")
			}
		case <-time.After(time.Second):
			fmt.Println("Timeout occurred, breaking loop")
			break
		case <-done:
			return
		}

		if aliceCh == nil && bobCh == nil {
			break
		}
	}
}
