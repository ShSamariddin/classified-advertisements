package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ShSamariddin/classified-advertisements/service"
)

func Auth(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	type AuthRequest struct {
		Phone string `json:"phone"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var ar AuthRequest
	_ = json.Unmarshal(body, &ar)
	_ = ar

	user, err := service.GetUser(ctx, ar.Phone)
	if err != nil {
		w.WriteHeader(404)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(200)
	_, _ = w.Write([]byte(*user.Firstname))
}
