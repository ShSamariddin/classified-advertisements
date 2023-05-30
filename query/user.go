package query

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/ShSamariddin/classified-advertisements/db"
	"github.com/ShSamariddin/classified-advertisements/model"
	"github.com/ShSamariddin/classified-advertisements/utils"
	"log"
)

func AddToTable(user model.User) {
	values, names := utils.GetStructFields(user, map[string]bool{"Id": true})
	fmt.Println(names)
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("users").
		Columns(names...).
		Values(values...).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Fatal("Error inserting user:", err)
	}

	var insertedID int64
	err = db.GetPool().QueryRow(context.Background(), query, args...).Scan(&insertedID)
	if err != nil {
		log.Fatal("Error executing insert query:", err)
	}
	fmt.Printf("insertedID: %d", insertedID)
}

func GetUserByPhone(phone string) model.User {
	fmt.Println(phone)
	query, args, err := sq.StatementBuilder.
		Select("*").
		From("users").
		Where(sq.Eq{"phone": phone}).
		ToSql()
	if err != nil {
		log.Fatal("Error selecting user by phone:", err)
	}
	user := model.User{}
	err = db.GetPool().QueryRow(context.Background(), query, args...).Scan(&user)
	fmt.Println(user)
	if err != nil {
		log.Fatal("Error executing select query:", err)
	}
	return user
}
