package shell

import (
	"fmt"
)

func mapActions(dependencies *Dependencies) {
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Write File")
		fmt.Println("2. Read File")
		fmt.Println("0. exit")

		var opcion int
		fmt.Scanln(&opcion)
		switch opcion {
		case 0:
			fmt.Println("going out...")
			return
		case 1:
			var location string
			fmt.Scanln(&location)
			answer := dependencies.writeHandler.WriteFile(location)
			fmt.Print(answer)
		case 2:
			var fileName string
			fmt.Scanln(&fileName)
			answer := dependencies.readHandler.ReadFile(fileName)
			fmt.Print(answer)
		case 3:
		default:
			fmt.Println("Invalid Option")
		}
	}
}
