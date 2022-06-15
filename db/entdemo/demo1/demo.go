package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"learn-go/db/entdemo/ent"
	"learn-go/db/entdemo/ent/education"
	"log"
)

func main() {
	client, err := ent.Open("mysql", "root:admin@tcp(localhost)/330test?charset=utf8mb4",
		ent.Debug())
	if err != nil {
		log.Fatalln(err)
	}
	_, err = client.Institution.Create().SetName("test institution").Save(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	edu, err := client.Education.Query().Where(education.ID(1)).Only(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(edu)
	all, err := edu.QueryInstitution().All(context.TODO())
	fmt.Println(all)

	client.Education.QueryInstitution(edu)
	//all, err := client.Education.QueryInstitution(client.Education.Query().Where(education.ID(1)).).All(context.TODO())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(all)
}
