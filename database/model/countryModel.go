package model

import (
	pb "../../proto"
	"database/sql"
	"fmt"
)

type Database struct {
	Db *sql.DB
}

func (database Database) Create(country *pb.Country) error {
	fmt.Println(country)
	result, err := database.Db.Exec("insert into country_schema.country(name, peopleNumber) values(?, ?)", country.Name, country.PeopleNumber)
	if err != nil {
		return err
	} else {
		country.Id, _ = result.LastInsertId()
	}
	return nil
}
func (database Database) Update(country *pb.Country) error {
	fmt.Println(country)
	_, err := database.Db.Exec("update country_schema.country set name = ?, peopleNumber = ?  where id = ?", country.Name, country.PeopleNumber, country.Id)
	if err != nil {
		return err
	}
	return nil
}
func (database Database) Delete(id int64) error {
	_, err := database.Db.Exec("delete from country_schema.country where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
func (database Database) Read(id int64) (pb.Country, error) {
	rows, err := database.Db.Query("select * from country_schema.country where id = ?", id)
	if err != nil {
		return pb.Country{}, err
	} else {
		var country pb.Country
		for rows.Next() {
			var id int64
			var name string
			var peopleNumber int64
			err := rows.Scan(&id, &name, &peopleNumber)
			if err != nil {
				return pb.Country{}, err
			} else {
				country = pb.Country{id, name, peopleNumber}
			}
		}
		return country, nil
	}
}
func (database Database) ReadAll() ([]*pb.Country, error) {
	rows, err := database.Db.Query("select * from country_schema.country")
	if err != nil {
		return nil, err
	} else {
		countries := []*pb.Country{}
		for rows.Next() {
			var id int64
			var name string
			var peopleNumber int64
			err := rows.Scan(&id, &name, &peopleNumber)
			if err != nil {
				return nil, err
			} else {
				user := pb.Country{id, name, peopleNumber}
				countries = append(countries, &user)
			}
		}
		return countries, nil
	}
}
