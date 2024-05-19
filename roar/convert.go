package roar

import "fmt"

func AsError(value any) error {
	if v, ok := value.(error); ok {
		return v
	}
	return fmt.Errorf("an error occurred: %v", value)
}
