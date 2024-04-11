package consoledemo

import "fmt"

func bla() {

}

// saker med små bokstäver är  INTERNA för packaget
func RunDemo() {
	var namn string
	fmt.Print("Ange namn:")
	fmt.Scanln(&namn)

	var age int
	fmt.Print("Ange ålder:")
	fmt.Scanln(&age)

	//fmt.Printf("Hej %s du är %d år\n", namn, age)
	fmt.Printf("Hej %v du är %v år\n", namn, age)

}
