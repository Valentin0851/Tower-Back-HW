package hw4

import "fmt"

func intSquare(n int, ans *int) {
	*ans += n * n
}

func doTask3() {
	nums := [...]int{2, 4, 6, 8, 10}
	ans := 0
	for _, val := range nums {
		go intSquare(val, &ans)
	}
	fmt.Println(ans)
	fmt.Scanln()
}
