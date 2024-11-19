package api

import (
	"fmt"
	"github.com/Solwery-Veronika/gateway/internal/model"
	"net/http"
	"strconv"
)

type Handler struct {
	srv SrvI
}

func New(srvs SrvI) *Handler { //конструктор
	return &Handler{
		srv: srvs,
	}
}

func (h *Handler) Handle(a, b int) int {
	data := model.Data{A: a, B: b}
	h.srv.Summator(&data) //Handler->srv->Summator
	return data.Sum
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		a := r.URL.Query().Get("a")
		b := r.URL.Query().Get("b")

		aInt, err := strconv.Atoi(a) //преобразуем в число
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		bInt, err := strconv.Atoi(b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		data := &model.Data{A: aInt, B: bInt}
		h.srv.Summator(data)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%d", data.Sum)))
	}
}
