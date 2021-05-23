package chat

import "fmt"

func Say(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
