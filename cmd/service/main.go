package main

import (
	"net/http"

	"github.com/Solwery-Veronika/gateway/internal/api"
	"github.com/Solwery-Veronika/gateway/internal/client"
	"github.com/Solwery-Veronika/gateway/internal/repository"
	"github.com/Solwery-Veronika/gateway/internal/service"
)

// компилируемые (go, c, c++, java)
// интерпритируемые (python, js)

func main() {
	repo := repository.New()

	// клиент
	c := client.New()

	srv := service.New(repo, c)

	handler := api.New(srv)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second) // контроль времени выполнения запроса
	//
	//defer cancel()

	//out, err := c.Login(ctx, &auth.LoginIn{
	//	Username: "Privet",
	//	Password: "1234",
	//})
	//if err != nil {
	//	log.Fatalf("error response: %v", err)
	//}

	// fmt.Println(out)

	http.HandleFunc("/signup", handler.Signup)

	// fmt.Println(res)

	err := http.ListenAndServe(":3232", nil)
	if err != nil {
		panic(err)
	}
}
