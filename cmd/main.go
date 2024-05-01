package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// admin "github.com/Abduazim0811/AvtoSalon/internal/Admin"
	sin "github.com/Abduazim0811/AvtoSalon/internal/Signin"
	sup "github.com/Abduazim0811/AvtoSalon/internal/Signup"
	_ "github.com/lib/pq"
)

func main() {
	var (
		num int
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Abdu0811", "dictionary")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	defer db.Close()


	dbnames := []string{"createuser", "createauto"}
	for _, dbname := range dbnames {
		name := "../internal/DB/" + dbname + ".sql"
		sqlfile, err := os.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}

		_,err=db.Exec(string(sqlfile))
		if err!=nil{
			log.Fatal(err)
		}
	}

	fmt.Println("[1]Users\t[2]Admin")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("[1]Signin\t[2]Signup\t[3]Exit")
		fmt.Scanln(&num)
		if num == 1 {
			sin.Signin(db)
		} else if num == 2 {
			sup.Signup(db)
		} else if num == 3 {
			os.Exit(0)
		}
	case 2:
		fmt.Println("Tez orada admin panel ishlaydi!!!")
		// admin.Admin()
	}
}
