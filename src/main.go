package main

import (
	"fmt"
	"net/http"
)

func main() {
	// rota que responde a requisições GET na raiz "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}