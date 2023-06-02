package service

import (
	"context"

	"github.com/ShSamariddin/classified-advertisements/model"
	"github.com/ShSamariddin/classified-advertisements/repository"
)

func GetUser(ctx context.Context, phone string) (model.User, error) {
	// TODO: regexp for checking phone number
	return repository.GetUserByPhone(ctx, phone)
}
