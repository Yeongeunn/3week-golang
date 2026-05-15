package main

import (
	"errors"
	"fmt"
)

func greet(name string) string {
	return "안녕하세요, " + name + "님!"
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	return a / b, nil
}

//두 정수를 받아 더한 값을 반환하는 `add` 함수를 만들고 `main`에서 호출해 보세요.
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greet("Alice"))

	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("result: %.2f\n", result)
	}

	// result, err = divide(5, 0)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }

	sum := add(3, 4)

	fmt.Println("sum:", sum)
}
