package signin

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	pr "github.com/Abduazim0811/AvtoSalon/internal/product"
	md "github.com/Abduazim0811/AvtoSalon/internal/models"
)

func Signin(db *sql.DB) {
	var (
		us md.User
		mp map[string]string=make(map[string]string)
	)
	ctx, cansel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cansel()
	fmt.Print("Email: ")
	fmt.Scanln(&us.Email)
	fmt.Print("Password: ")
	fmt.Scanln(&us.Password)
	name := "../internal/DB/insertuser.sql"
	query, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("Failed to read sql file: ", err)
	}

	rows, err := db.QueryContext(ctx, string(query))
	if err != nil {
		log.Fatal("Failed query: ", err)
	}

	for rows.Next() {
		var email, password string
		if err := rows.Scan(&email, &password); err != nil {
			log.Fatal("Failed")
		}
		mp[email] = password
	}

	value, found := mp[us.Email]
	if found {
		if strings.TrimSpace(us.Password) == strings.TrimSpace(value) {
			fmt.Println("Muvaffaqiyatli kirdingiz")
			pr.Product(db)
		}else{
			fmt.Println("Parol xato!!!")
		}
	}else{
		fmt.Println("Email xato!!!")
	}

}
