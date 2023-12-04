package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/factorial", HandlerFactorial)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans
}

func HandlerFactorial(w http.ResponseWriter, r *http.Request) {
	num := r.FormValue("num")
	n, err := strconv.Atoi(num)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	ans := factorial(n)
	w.Write([]byte(strconv.Itoa(ans)))
}
