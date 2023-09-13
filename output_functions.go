package the_first_module

import "fmt"

func OutputToConsole(txt string) error {
	_, err := fmt.Println(txt)
	return err
}
