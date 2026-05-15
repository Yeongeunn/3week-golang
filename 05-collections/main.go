package main

import "fmt"

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Println("array:", numbers)
	fmt.Println("first number:", numbers[0])

	names := []string{"Alice", "Bob", "Carol"}
	names = append(names, "Dave")

	//scores에 본인 이름과 점수를 추가하고, 전체 평균 점수를 계산해 보세요.

	fmt.Println("slice:", names)
	fmt.Println("slice length:", len(names))

	scores := map[string]int{
		"Alice": 90,
		"Bob":   75,
	}
	scores["Carol"] = 88


	//scores에 본인 이름과 점수를 추가
	scores["Youngeun"] = 95

	aliceScore, ok := scores["Alice"]
	if ok {
		fmt.Println("Alice score:", aliceScore)
	}

	total := 0 //전체 점수

	for name, score := range scores {
		fmt.Printf("%s: %d점\n", name, score)
		total += score
	}

	average := float64(total) / float64(len(scores))
	fmt.Printf("평균 점수: %.2f점", average)
}
