package elementary

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Elementary1() {
	fmt.Println("\nElementary-1")
	fmt.Println("Hello World")
}

func Elementary2() {
	fmt.Println("\nElementary-2")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type your good name please: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %v!\n", name)
}

func Elementary3() {
	fmt.Println("\nElementary-3")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please type your good name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "Alice" || name == "Bob" {
		fmt.Printf("Hello %v", name)
	} else {
		fmt.Printf("Sorry, only Alice and Bob are allowed :(\n")
	}
}

func Elementary4() {
	fmt.Println("\nElementary-4")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please type a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid number %v\n", input)
		return
	}
	result := 0
	for i := 1; i <= number; i++ {
		result += i
	}
	fmt.Printf("The sum from 1 upto the given number is %v\n", result)
}

func Elementary5() {
	fmt.Println("\nElementary-5")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please type a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid number %v\n", input)
		return
	}
	result := 0
	for i := 1; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	fmt.Printf("The sum from 1 upto %v (multiples of 3 and 5) is %v\n", number, result)
}

func Elementary6() {
	fmt.Println("\nElementary-6")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please type a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid number %v\n", input)
		return
	}

	actions := map[string]func(a, b int) int{
		"sum": func(a, b int) int { return a + b },
		"multiply": func(a, b int) int {
			if a == 0 {
				a = 1
			}
			return a * b
		},
	}
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Please type a action to perform: ")
	input, _ = reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))
	action, ok := actions[input]
	if !ok {
		fmt.Printf("Invalid action %v\n", input)
		return
	}
	result := 0
	for i := 1; i <= number; i++ {
		result = action(result, i)
	}
	fmt.Printf("The %v from 1 upto %v is %v\n", input, number, result)
}

func Elementary7() {
	fmt.Println("\nElementary-7")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please type a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid number %v\n", input)
		return
	}
	for i := 1; i <= 12; i++ {
		fmt.Printf("%v x %v = %v\n", number, i, number*i)
	}
}

func Elementary8() {
	fmt.Println("\nElementary-8")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Print prime numbers upto? (type a number): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid number %v\n", input)
		return
	}
	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	for i := 2; i <= number; i++ {
		if isPrime(i) {
			fmt.Printf("%v, ", i)
		}
	}
	fmt.Println("")
}

func Elementary9() {
	fmt.Println("\nElementary-9")
	secretNumber := 35
	attempts := map[int]int{}
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter your guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Invalid number %v\n", input)
			return
		}
		guessed, ok := attempts[guess]
		if ok {
			attempts[guess] = guessed + 1
		} else {
			attempts[guess] = 1
		}
		if guess < secretNumber {
			fmt.Println("The guess is too small. Please try again.")
		} else if guess > secretNumber {
			fmt.Println("The guess is too large. Please try again.")
		} else {
			fmt.Printf("Your guess is correct. You guessed in %v attempt(s).\n", len(attempts))
			break
		}
	}
}

func Elementary10() {
	fmt.Println("\nElementary-10")
	y := time.Now().Year() + 1
	counter := 1
	for counter <= 20 {
		if y%4 == 0 {
			fmt.Printf("%v: %v\n", counter, y)
			counter++
		}
		y++
	}
}

func Elementary11() {
	fmt.Println("\nElementary-11")
}
