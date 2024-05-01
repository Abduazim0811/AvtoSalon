package signup

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	pr "github.com/Abduazim0811/AvtoSalon/internal/product"
	md "github.com/Abduazim0811/AvtoSalon/internal/models"
)

func Signup(db *sql.DB) {
	var user md.User
	fmt.Print("Name: ")
	fmt.Scanln(&user.Name)
	fmt.Print("Email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Password: ")
	fmt.Scanln(&user.Password)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	// defer cancel()
	// chn:=make(chan int)

	// go func{

	// }()
	name := "../internal/DB/insertuser.sql"
	sqlfile, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	_, err = db.ExecContext(ctx, string(sqlfile), user.Name, user.Email, user.Password)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	fmt.Println("You have successfully signed up!")
	pr.Product(db)
	// go func() {
	// 	select {
	// 	case <-time.After(1* time.Second):
	// 		fmt.Println("Siz muvaffaqiyatli kirdingiz")
	// 	case <-ctx.Done():
	// 		fmt.Println("Muddat tugadi:", ctx.Err())
	// 	}
	// }()
}
