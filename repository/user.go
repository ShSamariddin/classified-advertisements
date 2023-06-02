package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/ShSamariddin/classified-advertisements/db"
	"github.com/ShSamariddin/classified-advertisements/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func GetUserByPhone(ctx context.Context, phone string) (model.User, error) {
	return getUserByColumn(ctx, "phone", phone)
}

func IsUserExist(ctx context.Context, id int64) (bool, error) {
	user, err := getUserByColumn(ctx, "id", id)
	if err != nil {
		return false, err
	}

	return user.ID == id, nil
}

func getUserByColumn(ctx context.Context, column string, value any) (model.User, error) {
	query := qb().Select("*").From("users").Where(sq.Eq{column: value})
	stmt, args, err := query.ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("unable to select from users: %v", err)
	}

	var user model.User
	err = pgxscan.Get(ctx, db.GetPool(), &user, stmt, args...)
	if err != nil {
		return model.User{}, fmt.Errorf("unable to get user: %v", err)
	}

	return user, nil
}
