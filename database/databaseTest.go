package main

import (
	"../database/config"
	"../database/model"
	pb "../proto"
	"fmt"
)

func main() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		countryModel := model.Database{
			Db: db,
		}
		country := pb.Country{
			Name:         "SomeName",
			PeopleNumber: 20,
		}
		err = countryModel.Create(&country)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("creating...")
			fmt.Println(country)
		}
		country = pb.Country{
			Id:           1,
			Name:         "SomeName2",
			PeopleNumber: 22,
		}
		err = countryModel.Update(&country)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("updating...")
		}
		err = countryModel.Delete(1)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("removing...")
		}
		country, err := countryModel.Read(2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("getting...")
			fmt.Println(country)
		}
		countries, err := countryModel.ReadAll()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("getting All...")
			fmt.Println(countries)
		}
	}

}
