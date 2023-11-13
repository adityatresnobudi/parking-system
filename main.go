package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adityatresnobudi/parking-system/parking"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func outputHandler(err error, a ...any) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(a...)
}

func main() {
	var attendant *parking.Attendant
	separator := "-------------------"
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	menu := "Parking Lot\n" +
		"1. Setup\n" +
		"2. Park\n" +
		"3. Un Park\n" +
		"4. Status\n" +
		"5. Exit"

	for !exit {
		fmt.Println(separator)
		fmt.Println(menu)
		input := promptInput(scanner, "input menu: ")

		switch input {
		case "1":
			capacities := promptInput(scanner, "input parking lot capacities: ")
			res, err := parking.SetupHandler(capacities)
			attendant = res
			outputHandler(err)
		case "2":
			plateNumber := promptInput(scanner, "input plate number: ")
			res, err := parking.ParkHandler(plateNumber, attendant)
			outputHandler(err, res)
		case "3":
			ticket := promptInput(scanner, "input ticket id: ")
			res, err := parking.UnParkHandler(ticket, attendant)
			outputHandler(err, res)
		case "4":
			res, err := parking.StatusHandler(attendant)
			outputHandler(err, res)
		case "5":
			exit = true
		default:
			fmt.Println("invalid menu")
		}
	}
}
