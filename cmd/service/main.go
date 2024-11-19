package main

import (
	"net/http"

	"github.com/Solwery-Veronika/gateway/internal/api"
	"github.com/Solwery-Veronika/gateway/internal/repository"
	"github.com/Solwery-Veronika/gateway/internal/service"
)

// компилируемые (go, c, c++, java)
// интерпритируемые (python, js)

func main() {
	repo := repository.New()

	srv := service.New(repo)

	handler := api.New(srv)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // ответ, запрос
		switch r.Method { // метод получения
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)   //статус 200 300 400 500
			w.Write([]byte("Hello GET!!")) //запись сообщения
		case http.MethodPost: //метод добавить
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Hello POST!!"))
		}
	})

	http.HandleFunc("/user/{id}", handler.Handler)

	//fmt.Println(res)

	err := http.ListenAndServe(":3232", nil)
	if err != nil {
		panic(err)
	}
}
