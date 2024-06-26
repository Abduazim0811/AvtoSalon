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
	fmt.Println("\nMashinalardan birini tanlang")

	name := "../internal/DB/selectauto.sql"

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
			year,id,count int
			price float64
		)

		if err:=rows.Scan(&id,&brand,&model,&year,&color,&price,&count); err!=nil{
			log.Fatal(err)
		}

		fmt.Printf("\nId: %d\nBrand: %s\nModel: %s\nYear: %d\nColor: %s\nPrice: %f\n",id,brand,model,year,color,price )
	}
	fmt.Println("Yuqoridagilardan birini tanlang!!/nSonini kiriting:")
	fmt.Scanln(&num)
	

	name="../internal/DB/updateauto.sql"
	sqlfile,err:=os.ReadFile(name)
	if err!=nil{
		log.Fatal(err)
	}
	_,err=db.Exec(string(sqlfile),num)
	if err!=nil{
		log.Fatal(err)
	}

}
