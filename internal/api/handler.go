package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Solwery-Veronika/gateway/internal/model"
)

type Handler struct {
	authService SrvI
}

func New(aS SrvI) *Handler { // конструктор
	return &Handler{
		authService: aS,
	}
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // если запрос идет не по методы отправки
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// конвертируем в массив байтов

	jsn, err := io.ReadAll(r.Body) // считываение байтов из тела запроса
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
			log.Printf("failed to write error response: %v", writeErr)
		} // сообщение об ошибке
		return
	}

	// преобразовываем массив байтов в структуру

	var data model.SignupData
	err = json.Unmarshal(jsn, &data) // превращение байтов в структуру(jsn-массив байтов, data-куда записать)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
			log.Printf("failed to write error response: %v", writeErr)
		}
		return
	}

	res, err := h.authService.SignupUsecase(r.Context(), data) // запускает процесс регистрации пользователя в сервисе аутентификации
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
			log.Printf("failed to write error response: %v", writeErr)
		}
		return
	}

	jsnBytes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
			log.Printf("failed to write error response: %v", writeErr)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsnBytes); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	jsn, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
			log.Printf("failed to write error response: %v", writeErr)
		}
		return
	}

	var data model.LoginData
	err = json.Unmarshal(jsn, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
			log.Printf("failed to write error response: %v", writeErr)
		}
		return
	}

	res, err := h.authService.LoginUsecase(r.Context(), data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
			log.Printf("failed to write error response: %v", writeErr)
		}
		return
	}

	jsnBytes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
			log.Printf("failed to write error response: %v", writeErr)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsnBytes); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
