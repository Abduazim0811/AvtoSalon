package product

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func Product(db *sql.DB) {
	var (
		num int
	)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	fmt.Println("Mashinalardan birini tanlang")

	name := "..internal/DB/selectauto.sql"

	query, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("File to read sql file ", err)
	}

	rows, err := db.QueryContext(ctx, string(query))
	if err!=nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var (
			brand,model,color string
			year,id int
			price float64
		)

		if err:=rows.Scan(&id,&brand,&model,&year,&color,&price); err!=nil{
			log.Fatal(err)
		}

		fmt.Printf("Id: %d\nBrand: %s\nModel: %s\n, Year: %d\n,Color: %s\n,Price: %f\n",id,brand,model,year,color,price )
	}
	fmt.Println("Yuqoridagilardan birini tanlang!!/nSonini kiriting:")
	fmt.Scanln(&num)
	name="../internal/DB/selectwhereauto.sql"

	sqlfile,err:=os.ReadFile(name)
	if err!=nil{
		log.Fatal(err)
	}
	
	_,err=db.Query(string(sqlfile),num)
	if err!=nil{
		log.Fatal("Error id")
	}

	name="../internal/DB/updateauto.sql"
	sqlfile,err=os.ReadFile(name)
	if err!=nil{
		log.Fatal(err)
	}
	_,err=db.Query(string(sqlfile),num)
	if err!=nil{
		log.Fatal(err)
	}

}
