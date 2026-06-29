// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package tools

import (
	"os"
)

// this function maded for clear user's screen(terminal). speed and easy =D
func ClearScreen() {
	os.Stdout.WriteString("\033[2J\033[3J\033[H") // write this code to clear terminal page. i use os.Stdout.WriteString for speedly.
}
