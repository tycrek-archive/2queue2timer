package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"

	"github.com/go-vgo/robotgo"
)

/* Strings */
var pressEnter = "\nPress Enter to continue..."
var done = "\n\nLiftoff! Joining 2b2t queue\n"
var starting = "Starting countdown (you may not see any output for up to 60 seconds)\n"

var step_1 = "\nStep 1: Capture mouse pointer"
var step_1_1 ="  - Once you press Enter, you have 3 seconds to position your mouse pointer over the play button for 2b2t"
var step_1_2 ="  - Have Minecraft in windowed mode to prevent any problems"
var step_1_3 ="  - To cancel at any point in time, press CTRL+C"


var step_2 = "\n\nStep 2: Enter queue time"
var step_2_1 = "  - Enter a time for 2q2t to join the queue"
var step_2_2 = "  - At this time, 2q2t will auto click the play button to join the queue"
var step_2_3 = "  - Use 24-hour time (example: 10 PM = 22)"
var step_2_4 = "  - Use the following format: HH:MM"

var step_3 = "\n\nStep 3: Wait"
var step_3_1 = "  - When this countdown reaches 0, 2q2t will click the play button"
var step_3_2 = "  - Make sure the play button is visible"
var step_3_3 = "  - Do not move the Minecraft window from where it is now until you have joined queue"
var step_3_4 = "  - Disable automatic sleep or lock on your PC"

/* Variables */
var x int
var y int
var userTime string

/* Functions */

func formatRemaining(h int, m int, s int) {
	fmt.Print("\r> ", h, " hours ", m, " minutes ", s," seconds left        ")
}

// Prints the remaining time until click
func printRemaining(remaining int) {
	hLeft := int(remaining / (60 * 60))
	mLeft := int(remaining / 60 % 60)
	sLeft := int(remaining % 60)

	formatRemaining(hLeft,mLeft,sLeft)
}

// Clicks the mouse
func click() {
	fmt.Println(done)
	robotgo.MoveMouse(x, y)
	robotgo.Click()
}

// Print welcome message
func welcome() {
	fmt.Println("\n")
	fmt.Println("######################################")
	fmt.Println("##                                  ##")
	fmt.Println("##         ~ 2queue2timer ~         ##")
	fmt.Println("##                                  ##")
	fmt.Println("##         Author:  tycrek          ##")
	fmt.Println("##        Website:  jmoore.dev      ##")
	fmt.Println("##                                  ##")
	fmt.Println("######################################")
	fmt.Println("\n")
}

// Step 1: record mouse position
func step1() {
	fmt.Println(step_1)
	fmt.Println(step_1_1)
	fmt.Println(step_1_2)
	fmt.Println(step_1_3)
	fmt.Println(pressEnter)
	fmt.Scanln()

	// Wait 3 seconds and record mouse position
	time.Sleep(3 * time.Second)
	x, y = robotgo.GetMousePos()
	fmt.Println("Mouse recorded at: [", x, "] [", y, "]")
}

// Step 2: Get time to queue from user
func step2() {
	fmt.Println(step_2)
	fmt.Println(step_2_1)
	fmt.Println(step_2_2)
	fmt.Println(step_2_3)
	fmt.Println(step_2_4)

	fmt.Print("> ")
	fmt.Scanln(&userTime)
	fmt.Println("\nWill queue at: ", userTime)
}

// Step 3: Count down and click play
func step3() {
	fmt.Println(step_3)
	fmt.Println(step_3_1)
	fmt.Println(step_3_2)
	fmt.Println(step_3_3)
	fmt.Println(step_3_4)
	fmt.Println(pressEnter)
	fmt.Scanln()
	fmt.Println(starting)

	// Split the user entered string to get the hour and minute components
	qtSplit := strings.Split(userTime, ":")
	qtHour, _ := strconv.Atoi(qtSplit[0])
	qtMinute, _ := strconv.Atoi(qtSplit[1])
	
	// Get the current time and the users timezone
	t := time.Now()
	tzone, offset := t.Zone()
	loc :=  time.FixedZone(tzone, offset)

	// Check if time is for the next day
	qtDay := t.Day()
	if qtHour < t.Hour() {
		qtDay++
	}

	// Creat a Date object for the queue time
	qt := time.Date(t.Year(), t.Month(), qtDay, qtHour, qtMinute, 0, 0, loc)

	// Loop until it is time to queue
	for {
		remaining := int(time.Until(qt).Seconds())
		if remaining > 0 {
			printRemaining(remaining)
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	click()
}

/* Main */
func main() {
	welcome()
	step1()
	step2()
	step3()
}