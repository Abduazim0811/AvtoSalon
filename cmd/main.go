package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	admin "github.com/Abduazim0811/AvtoSalon/internal/Admin"
	sin "github.com/Abduazim0811/AvtoSalon/internal/Signin"
	sup "github.com/Abduazim0811/AvtoSalon/internal/Signup"
)

func main() {
	var (
		num int
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "localhost", "Abdu0811", "dictionary")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("[1]Users\t[2]Admin")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("[1]Signin\t[2]Signup\t[3]Exit")
		fmt.Scanln(&num)
		if num == 1 {
			sin.Signin()
		} else if num == 2 {
			sup.Signup()
		} else if num == 3 {
			os.Exit(0)
		}
	case 2:
		admin.Admin()
	}
}
