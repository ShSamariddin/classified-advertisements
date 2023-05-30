package main

import (
	"fmt"
	"github.com/ShSamariddin/classified-advertisements/db"
	"github.com/ShSamariddin/classified-advertisements/query"
)

func main() {
	//user := model.User{Firstname: "samar", Lastname: "sharipov", Phone: "+79692005003"}
	pool := db.GetPool()

	defer pool.Close()
	//query.AddToTable(user)
	selectUser := query.GetUserByPhone("+79692005003")
	fmt.Println(selectUser)
}
