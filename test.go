// package main

// import "fmt"

// func main() {
// 	color1, color2, color3 := colors()

// 	fmt.Println(color1, color2, color3)
// }

// func colors() (string, string, string) {
// 	return "red", "yellow", "blue"
// }

package main

import "fmt"

func main() {
	c := color("Red")

	fmt.Println(c.describe("is an awesome color"))
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}
