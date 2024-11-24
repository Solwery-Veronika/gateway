package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Solwery-Veronika/gateway/internal/api"
	"github.com/Solwery-Veronika/gateway/internal/client"
	"github.com/Solwery-Veronika/gateway/internal/repository"
	"github.com/Solwery-Veronika/gateway/internal/service"
	"github.com/Solwery-Veronika/gateway/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// компилируемые (go, c, c++, java)
// интерпритируемые (python, js)

func main() {
	repo := repository.New()

	srv := service.New(repo)

	handler := api.New(srv)

	// клиент

	conn, err := grpc.NewClient("0.0.0.0:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // закрыть соединение

	grpcClient := auth.NewAuthServiceClient(conn) // создаем наш сгенерированый клиент

	c := client.New(grpcClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // контроль времени выполнения запроса

	defer cancel()

	out, err := c.Login(ctx, &auth.LoginIn{
		Username: "Privet",
		Password: "1234",
	})
	if err != nil {
		log.Fatalf("error response: %v", err)
	}

	fmt.Println(out)

	//-------------------------------------------

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // ответ, запрос
		switch r.Method { // метод получения
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)   // статус 200 300 400 500
			w.Write([]byte("Hello GET!!")) // запись сообщения
		case http.MethodPost: // метод добавить
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Hello POST!!"))
		}
	})

	http.HandleFunc("/user/{id}", handler.Handler)

	// fmt.Println(res)

	err = http.ListenAndServe(":3232", nil)
	if err != nil {
		panic(err)
	}
}
