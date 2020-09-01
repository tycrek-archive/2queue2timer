package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("\nWelcome to 2queue2timer")
	fmt.Println("\nAfter pressing Enter, move your mouse pointer to the Play button in the server list for 2b2t.")
	fmt.Println("After 3 seconds, 2q2t will capture that position.")
	fmt.Println("Press Enter to continue...")

	fmt.Scanln()

	time.Sleep(3 * time.Second)
	
	x, y := robotgo.GetMousePos()
	fmt.Println("Mouse recorded at: ", x, y)

	fmt.Println("\nNext enter a time to queue. Follow this format:")
	fmt.Println("    HH:MM")
	fmt.Println("Make sure to use 24 hour time (i.e. 10 PM = 22)")

	var qts string
	fmt.Print("> ")
	fmt.Scanln(&qts)

	fmt.Println("Will queue at:", qts, "Make sure you have sleep and lock DISABLED on your PC. Also do not move the Minecraft window until the program has joined queue.")

	qtsplit := strings.Split(qts, ":")
	qthour, _ := strconv.Atoi(qtsplit[0])
	qtminute, _ := strconv.Atoi(qtsplit[1])

	t := time.Now()
	tzone, offset := t.Zone()
	loc :=  time.FixedZone(tzone, offset)

	qt := time.Date(t.Year(), t.Month(), t.Day(), qthour, qtminute, 0, 0, loc)
	
	for {
		if int(time.Until(qt).Seconds()) > 1 {
			fmt.Println(int(time.Until(qt).Seconds()), " seconds left")
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	fmt.Println("yeet")

	robotgo.MoveMouse(x, y)
	robotgo.Click()

}