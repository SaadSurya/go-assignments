package list_and_strings

import (
	"fmt"
)

func ListAndStrings1(list []int) int {
	fmt.Println("\nList & Strings 1")
	max := 0
	for _, n := range list {
		if max < n {
			max = n
		}
	}
	fmt.Printf("Largest number in the list is: %v\n", max)
	return max
}

func ListAndStrings2(list []int) []int {
	fmt.Println("\nList & Strings 2")
	for i := 0; i < len(list)/2; i++ {
		j := len(list)-1-i
		if i != j {
			temp := list[i]
			list[i] = list[j]
			list[j] = temp
		}
	}
	fmt.Printf("The reversed list is: %v\n", list)
	return list
}

func ListAndStrings3(list []int, element int) bool {
	fmt.Println("\nList & Strings 3")
	for _, v := range list {
		if(v == element) {
			fmt.Printf("The element %v exists in the list.\n", element)
			return true			
		}
	}
	fmt.Printf("The element %v does not exists in the list.\n", element)
	return false
}

func ListAndStrings4(list []int) []int {
	fmt.Println("\nList & Strings 4")
	odds := make([]int, len(list)/2, (len(list)/2)+1)
	j := 0
	for i, v := range list {
		if (i+1)%2 == 1 {
			if len(odds)-j == 0 {
				odds = odds[:cap(odds)]
			}
			odds[j] = v
			j++
		}
	}
	fmt.Printf("Odd positioned elements are: %v\n", odds)
	return odds
}

func ListAndStrings5(list []int) []int {
	fmt.Println("\nList & Strings 5")
	rt := make([]int, len(list))
	for i, v := range list {
		if i == 0 {
			rt[i] = v
		} else {
			rt[i] = rt[i-1] + v
		}
	}
	fmt.Printf("Running totals of the given list are: %v\n", rt)
	return rt
}

func ListAndStrings6(s string) bool {
	fmt.Println("\nList & Strings 6")
	var r string
	b :=  true
	for i, v := range s {
		if v != rune(s[len(s)-1-i]) {
			r = " not"
			b = false
			break
		}
	}
	fmt.Printf("The string is%v palindrome.\n", r)
	return b
}

func ListAndStrings7(list []int) {
	fmt.Println("\nList & Strings 7")
	forResult := 0
	for _, v := range list {
		forResult += v
	}
	
	i := 0
	whileResult := 0
	for i < len(list) {
		whileResult += list[i]
		i++
	}

	recursionResult := sum(list, 0)
	fmt.Printf("forResult: %v, whileResult: %v, recursionResult: %v\n", forResult, whileResult, recursionResult)
}
func sum(list []int, i int) int {
	if(i < len(list)) {
		return 0
	}
	return sum(list, i+1) + list[i] 
}