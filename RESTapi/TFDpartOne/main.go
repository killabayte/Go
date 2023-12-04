package main

import (
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/factorial", HandlerFactorial)
	http.ListenAndServe(":8080", nil)
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
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}
	w.Write([]byte(strconv.Itoa(factorial(n))))
}
