package person

import "fmt"

func Description(name string) string {
	return fmt.Sprintf("The person name is %s", name)
}

func secretName(name string) string {
	return "Do not share."
}
