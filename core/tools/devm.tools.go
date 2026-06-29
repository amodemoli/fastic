// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package tools

import (
	"fmt"

	"github.com/amodemoli/fastic/core/color"
)

// this function maded for send a message on terminal if development mode is true
// you can change message symbol color, symbol character message model and message body.
func DevMessage(devM bool, clr, symbol, model, message string) {
	// if development mode is true print the message
	if devM {
		fmt.Printf("    %s%s%s %s: %s\n", clr, symbol, color.Nc, model, message) // send a message if evelopment mode is true
	}
}
