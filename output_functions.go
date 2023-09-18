package the_first_module

import (
	"fmt"
	"strings"
)

var TxtSeparatorChar string = "-"

func OutputToConsole(txt string) error {
	_, err := fmt.Println(txt)
	return err
}

func OutputSeparator() error {
	err := OutputToConsole(strings.Repeat(TxtSeparatorChar, 30))
	return err
}

func OutputToConsoleWithSeparator(txt string) error {
	err := OutputToConsole(txt)
	if err != nil {
		return err
	}
	return OutputSeparator()
}
