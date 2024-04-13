package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var input string
	ClearTerminal()

	var word = []string{"p", "i", "z", "z", "a"}
	var revealed = []string{"_", "_", "_", "_", "_"}
	var lives int8 = 5
	fmt.Printf("Lives: %d\n", lives)
	fmt.Println(strings.Join(revealed, ""))

	for {
		scanner.Scan()
		input = scanner.Text()
		//fmt.Println(input)
		ClearTerminal()

		if len(input) > 1 {
			fmt.Println("Only 1 charecter is expected")
			continue
		}

		if input == "_" {
			fmt.Println("The underscore is not allowed")
			continue
		}

		doesWordContainInput, arrayOfIndexs := doesArrayContainThis(input, word)
		if doesWordContainInput {
			for i := 0; i < len(arrayOfIndexs); i++ {
				revealed[arrayOfIndexs[i]] = string(input[0])
			}
		} else {
			lives--
			if lives == 0 {
				fmt.Println("You lost!")
				os.Exit(0)
			}
		}

		// for i := 0; i < len(word); i++ {
		// 	if word[i] == input[0] {
		// 		revealed[i] = string(input[0])
		// 	} else {
		// 		lives--
		// 		fmt.Println(lives)
		// 		if lives == 0 {
		// 			fmt.Println("You lost!")
		// 			os.Exit(0)
		// 		}
		// 		break
		// 	}
		// }

		fmt.Printf("Lives: %d\n", lives)
		fmt.Println(strings.Join(revealed, ""))

		if strings.Join(word, "") == strings.Join(revealed, "") {
			fmt.Println("Word found!")
			os.Exit(0)
		}
	}
}

func doesArrayContainThis(val string, stringArray []string) (bool, []int) {
	var foundLocations []int
	for i := 0; i < len(stringArray); i++ {
		if stringArray[i] == val {
			foundLocations = append(foundLocations, i)
		}
	}

	if len(foundLocations) == 0 {
		return false, foundLocations
	} else {
		return true, foundLocations
	}
}
