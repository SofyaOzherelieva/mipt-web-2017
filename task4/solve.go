package main

import (
	"fmt"
	"unicode"
)

func RemoveEven(a []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(a); i++{
		if a[i] % 2 == 1 {
			res = append(res, a[i])
		}
	}
	return res
}

func PowerGenerator(a int) (func() int) {
	i1 := 1
	return func() (power int) {
		power = a*i1
		i1 = power
		return
	}
}

func DifferentWordsCount(input_string string) int {
	word := ""
	answer := 0
	alreadyCount := make(map[string]bool)
	input_string = input_string + " "
	for _, symbol := range (input_string){
		if unicode.IsLetter(symbol) {
			word += string(unicode.ToLower(symbol))
		} else if word != "" {
			if alreadyCount[word] == false{
				answer += 1
				alreadyCount[word] = true
			}
			word = ""
		}
	}
	return answer
}


/*func main() {
	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
}*/
