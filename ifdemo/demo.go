package ifdemo

import "fmt"

func Test() {
	var namn = "Stefan"
	var y = -1

	if y < 0 {
		fmt.Printf("%d is negative\n", y)
	} else {
		fmt.Printf("%d is not  negative\n", y)
	}

	var age = 18
	if age >= 65 {
		fmt.Println("Du är pensionär")
	} else if age >= 18 {
		fmt.Println("Du är vuxen")
	} else {
		fmt.Println("Du är barn")
	}

	if namn == "Stefan" && age >= 49 {
		fmt.Println("Du är fortfarande ung")
	}
}
