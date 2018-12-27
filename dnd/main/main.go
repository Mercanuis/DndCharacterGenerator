package main

import (
	"bufio"
	"dnd/charsheet"
	"fmt"
	"os"
)

func main ()  {
	fmt.Println("Enter a name for your character")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sheet := charsheet.CreateNewCharacter(scanner.Text())
		charsheet.PrintSheet(sheet)
		break
	}
}