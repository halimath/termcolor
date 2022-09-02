package termcolor_test

import "github.com/halimath/termcolor"

func Example() {
	p := termcolor.Stdout()

	p.Printf("Welcome to %s output!", p.Styled("colored", termcolor.ForegroundCyan), termcolor.Bold)

	// Output: Welcome to colored output!
}
