package main

import (
	"fmt"
	"net/http"
)

// 1. SUM
func sum(w http.ResponseWriter, r *http.Request) {
	a := 50
	b := 50
	result := a + b

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"result": %d}`, result)
}

// 2. SUBTRACTION
func subtract(w http.ResponseWriter, r *http.Request) {
	a := 80
	b := 30
	result := a - b

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"result": %d}`, result)
}

// 3. MULTIPLICATION
func multiply(w http.ResponseWriter, r *http.Request) {
	a := 10
	b := 20
	result := a * b

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"result": %d}`, result)
}

// 4. DIVISION
func divide(w http.ResponseWriter, r *http.Request) {
	a := 100
	b := 20
	result := a / b

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"result": %d}`, result)
}

// 5. SQUARE
func square(w http.ResponseWriter, r *http.Request) {
	n := 5
	result := n * n

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"square": %d}`, result)
}

// 6. EVEN / ODD
func evenOdd(w http.ResponseWriter, r *http.Request) {
	n := 10

	result := "odd"

	if n%2 == 0 {
		result = "even"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"number": %d, "result": "%s"}`, n, result)
}

// 7. POSITIVE / NEGATIVE / ZERO
func checkNumber(w http.ResponseWriter, r *http.Request) {
	n := -5
	result := ""

	if n > 0 {
		result = "positive"
	} else if n < 0 {
		result = "negative"
	} else {
		result = "zero"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"number": %d, "result": "%s"}`, n, result)
}

// 8. MAXIMUM OF TWO NUMBERS
func maximum(w http.ResponseWriter, r *http.Request) {
	a := 50
	b := 80

	max := a

	if b > max {
		max = b
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"maximum": %d}`, max)
}

// 9. MINIMUM OF TWO NUMBERS
func minimum(w http.ResponseWriter, r *http.Request) {
	a := 50
	b := 80

	min := a

	if b < min {
		min = b
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"minimum": %d}`, min)
}

// 10. PRIME NUMBER CHECK
func prime(w http.ResponseWriter, r *http.Request) {
	n := 17
	isPrime := true

	if n < 2 {
		isPrime = false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"number": %d, "prime": %t}`, n, isPrime)
}

// 11. FACTORIAL
func factorial(w http.ResponseWriter, r *http.Request) {
	n := 5
	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"number": %d, "factorial": %d}`, n, result)
}

// 12. FIBONACCI
func fibonacci(w http.ResponseWriter, r *http.Request) {
	n := 10

	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"n": %d, "fibonacci": %d}`, n, a)
}

// 13. GCD
func gcd(w http.ResponseWriter, r *http.Request) {
	a := 24
	b := 36

	for b != 0 {
		a, b = b, a%b
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"gcd": %d}`, a)
}

// 14. PALINDROME
func palindrome(w http.ResponseWriter, r *http.Request) {
	n := 121
	original := n
	reverse := 0

	for n > 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n = n / 10
	}

	isPalindrome := original == reverse

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"number": %d, "palindrome": %t}`, original, isPalindrome)
}

// 15. LEAP YEAR
func leapYear(w http.ResponseWriter, r *http.Request) {
	year := 2024

	isLeap := false

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		isLeap = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"year": %d, "leap_year": %t}`, year, isLeap)
}

// 16. HEALTH CHECK
func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, `{"status":"okay","status_code":200}`)
}

func main() {

	http.HandleFunc("/sum", sum)
	http.HandleFunc("/subtract", subtract)
	http.HandleFunc("/multiply", multiply)
	http.HandleFunc("/divide", divide)
	http.HandleFunc("/square", square)
	http.HandleFunc("/even-odd", evenOdd)
	http.HandleFunc("/check-number", checkNumber)
	http.HandleFunc("/maximum", maximum)
	http.HandleFunc("/minimum", minimum)
	http.HandleFunc("/prime", prime)
	http.HandleFunc("/factorial", factorial)
	http.HandleFunc("/fibonacci", fibonacci)
	http.HandleFunc("/gcd", gcd)
	http.HandleFunc("/palindrome", palindrome)
	http.HandleFunc("/leap-year", leapYear)
	http.HandleFunc("/health", health)

	fmt.Println("Server started on port 8080")

	http.ListenAndServe(":8080", nil)
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	w.WriteHeader(200)

// 	fmt.Fprintln(w, `{"status":"ok","code":200}`)
// }

// func main() {

// 	http.HandleFunc("/health", healthCheckHandler)

// 	fmt.Println("Server started on port 8080")

// 	http.ListenAndServe(":8080", nil)
// }


// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func health(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprint(w, `{"status": "okay", "status_code": 200}`)
// }

// func main() {
// 	http.HandleFunc("/health", health)
// 	fmt.Println("Server started successfully!")
// 	fmt.Println("Sum API:    http://localhost:8080/sum")
// 	http.ListenAndServe(":8080", nil)
// }
