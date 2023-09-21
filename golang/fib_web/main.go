package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8089", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	res := fib(40)
	name, _ := os.Hostname()
	fmt.Fprintln(w, fmt.Sprintf("%d + %s", res, name))
}

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
