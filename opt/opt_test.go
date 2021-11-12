package opt_test

import (
	"fmt"
	"github.com/wwalexander/gofunc/opt"
)

func ExampleOpt() {
	hackers := []struct {
		Name   string
		Handle opt.Opt[string]
	}{
		{"Dade Murphy", opt.Some("Crash Override")},
		{"Kate Libby", opt.Some("Acid Burn")},
		{"Joey Pardella", opt.None[string]()},
	}
	for _, hacker := range hackers {
		fmt.Printf("%s: ", hacker.Name)
		identity := opt.Then(hacker.Handle, func(handle string) string {
			return fmt.Sprintf("I'm %s.", handle)
		})
		if introduction, ok := identity.Let(); ok {
			fmt.Println(introduction)
		} else {
			fmt.Println("I don't have an identity until I have a handle.")
		}
	}
	// Output:
	// Dade Murphy: I'm Crash Override.
	// Kate Libby: I'm Acid Burn.
	// Joey Pardella: I don't have an identity until I have a handle.
}
