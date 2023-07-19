package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	var now time.Time = time.Now()
// 	var year int = now.Year()
// 	fmt.Println(year)
// }

func main() {
	brocken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(brocken)
	fmt.Println(fixed)
}
