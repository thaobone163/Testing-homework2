package main

import "fmt"

func SalaryRank(sense, applicability, time int) string {
	if sense < 0 || applicability < 0 || time < 0 || sense > 100 || applicability > 100 || time > 100 {
		return "Input không hợp lệ"
	}
	total := sense + applicability*2 + time
	if sense < 90 || applicability < 45 || time < 70 || total < 260 {
		return "Rank C"
	} else {
		if applicability < 70 || total < 320 {
			return "Rank B"
		}
	}
	return "Rank A"
}

func main() {
	sense := 90
	applicability := 90
	time := 90
	rank := SalaryRank(sense, applicability, time)

	fmt.Printf("Result: %q\n", rank)
}
