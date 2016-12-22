// Package greeting implements the solution to the Hello World
// problem at http://exercism.io/exercises/go/hello-world/readme
package greeting

const testVersion = 3

// HelloWorld returns a string that greets the user by name, or says "Hello, World!" if no
// name is given
func HelloWorld(name string) string {
	var ret string

	if len(name) <= 0 {
		ret = "World"
	} else {
		ret = name
	}

	return "Hello, " + ret + "!"
}
