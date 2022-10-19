package main

// package = project = workspace
// Executable package = generates a file that we can run.
// Reusable package = code used as "helpers". Good place to put reusable logic. Aka code dependecies or libraries.
// How we know which kind are we making? By the name. For example: main is executable package. Anything else is a reusable.
// In package main there must be always a function called main.

import "fmt"

/*
If there are multiple lines for importing:
import (
	"fmt"
	"strings"
)

Also import mean: give my package <package name> access to all the code and functionality
that is contained inside of this other package called <imported package name>.

Standard library can be found on: www.golang.org/pkg
For example:
fmt = standartized format package
*/

func main() {
	fmt.Println("Hello world!")
}
