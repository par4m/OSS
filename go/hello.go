
import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello Arch")
	fmt.Println("Hi from Go")
	x := add(1, 2)
	fmt.Println(x)

	q, r := divide(5, 3)
	fmt.Println("WHen divided by 5 ", "quotient is ", q, "remainder is ", r)
	fmt.Print("Sum is ", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println()

	c := average(1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.0)
	fmt.Println(c)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Sum of nums is ", sum(nums...))

	fmt.Println("Concatenated string is ", concatenateStrings("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"))
	T()
}

func add(x, y int) int {
	return x + y
}

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Factorial is undefined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	fact, err := factorial(n - 1)
	if err != nil {
		return 0, err
	}
	return n * fact, nil
}

func divide(dividend, divisor int) (quotient int, remainder int) {
	if divisor == 0 {
		return
	}
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

func a1() {
	fmt.Println("new")
}

// variadic functions

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total

}

func average(nums ...float64) float64 {

	total := 0.0
	count := 0

	for _, num := range nums {
		total += num
		count++
	}
	if count == 0 {

		return 0
	}
	return total / float64(count)

}

func concatenateStrings(strings ...string) string {
	result := ""
	for _, s := range strings {
		result += s
	}
	return result

}
