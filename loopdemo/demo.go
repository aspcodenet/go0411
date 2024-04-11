package loopdemo

import "fmt"

func RunDemo() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)
}

func RunDemoWhile() {
	n := 1
	for n < 5 {
		fmt.Println(n)
		n++
	}
}

func RunDemoForeverLoop() {
	// while(true)
	for {
		fmt.Println("Go is weird")
	}
}

func RunDemoForEach() {
	// arrayer vs vector<>  vs List<>
	// arrayer eller slices

	allaNamn := []string{"Stefan", "Kalle", "Josefine"}

	// for each
	// for i, oneName := range allaNamn {
	// 	fmt.Printf("På plats %v finns %v \n", i, oneName)
	// }
	for _, oneName := range allaNamn {
		fmt.Printf("Namn: %v \n", oneName)
	}

}
