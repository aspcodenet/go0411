package funktioner

import "fmt"

// func add(x int, y int) {
// 	total := 0
// 	total = x + y
// 	fmt.Println(total)
// }

func areaRectangle2(l int, b int) int {
	var area int
	area = l * b
	return area
}

func areaRectangle(l int, b int) (area int, errorText string) {
	errorText = ""
	if l <= 0 {
		errorText = "Jamen det är negativ längd"
	}
	if b <= 0 {
		errorText = "Jamen det är negativ bredd"
	}
	area = l * b
	return
}

func Test() {
	area, errorTexten := areaRectangle(12, -1)
	if len(errorTexten) > 0 {
		fmt.Println("Fel")
	} else {
		fmt.Println(area)
	}

}

func Withdraw(currentBalance int, belopp int) (newBalance int, errorText string) {
	errorText = ""
	newBalance = currentBalance
	if belopp > currentBalance {
		errorText = "Så mycket pengar har du inte"
	} else if belopp > 3000 {
		errorText = "Maxbelopp 3000 per uttag"
	} else {
		newBalance = currentBalance - belopp
	}
	return
}
func Hej() {
	balance := 1000
	belopp := 1300
	errorText := ""
	balance, errorText = Withdraw(balance, belopp)
	if len(errorText) > 0 {
		fmt.Println(errorText)
	}
}
