package main

import (
	"fmt"
	"net/http"
)

// HelloWorld — обработчик запроса.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World</h1>"))
	fmt.Println(r.URL.Path)
}

func main() {
	// маршрутизация запросов обработчику
	http.HandleFunc("/", HelloWorld)
	// запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)
}
