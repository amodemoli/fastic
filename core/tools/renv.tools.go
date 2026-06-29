// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package tools

import (
	"os"
	"strings"
)

// get list of values from env file. this function get key and return value of ".env" file
// and this function supports default value, means if value is nil use default values =D
func ReadEnvValue(key, defaultValue string) string {
	val := strings.TrimSpace(os.Getenv(key)) // remove prefix/suffix spaces from response.
	// if key value is nil (i mean ""), they change value to default value's
	if val == "" {
		return defaultValue
	}
	// return value for caller =D with string's to lower
	return strings.ToLower(val)
}
