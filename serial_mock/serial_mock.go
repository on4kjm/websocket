package serial_mock

import (
	"fmt"
	"time"
)

var MessageBuffer = make(chan string, 10)

func Receive_loop() {
	ticker := time.NewTicker(5 * time.Second)

	i := 0
	// every time our ticker ticks
	for t := range ticker.C {
		output := fmt.Sprintf(t.Format("15:04:05 "))
		i++
		if i == 4 {
			output = output + "\n"
			i = 0
		}
		fmt.Printf("Serial mock: %s\n", output)

		//Send output to the websocket go-routine
		MessageBuffer <- output
	}
}
