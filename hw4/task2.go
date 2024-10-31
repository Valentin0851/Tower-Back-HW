package hw4

import "fmt"

func square(n int) {
	fmt.Println(n * n)
}

func doTask2() {
	nums := [...]int{2, 4, 6, 8, 10}

	for _, val := range nums {
		go square(val)
	}
	fmt.Scanln()
}
