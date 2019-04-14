package implementations

import (
	. "../../database/connection"
	"../../database/queries"
	. "../entities"
)

type CountryDAOImpl struct {
}

func (database CountryDAOImpl) Read(id int64) (*Country, error) {
	db := GetConnection()
	defer db.Close()
	rows, err := db.Query(queries.GetCountryById, id)
	if err != nil {
		return &Country{}, err
	} else {
		var country Country
		for rows.Next() {
			var id int64
			var name string
			var peopleNumber int64
			err := rows.Scan(&id, &name, &peopleNumber)
			if err != nil {
				return &Country{}, err
			} else {
				country = Country{Id: id, Name: name, PeopleNumber: peopleNumber}
			}
		}
		return &country, nil
	}
}

func (database CountryDAOImpl) ReadAll() ([]*Country, error) {
	db := GetConnection()
	defer db.Close()
	rows, err := db.Query(queries.GetCountries)
	if err != nil {
		return nil, err
	} else {
		var countries []*Country
		for rows.Next() {
			var id int64
			var name string
			var peopleNumber int64
			err := rows.Scan(&id, &name, &peopleNumber)
			if err != nil {
				return nil, err
			} else {
				user := Country{Id: id, Name: name, PeopleNumber: peopleNumber}
				countries = append(countries, &user)
			}
		}
		return countries, nil
	}
}

func (database CountryDAOImpl) Create(country *Country) (int64, error) {
	db := GetConnection()
	defer db.Close()
	result, err := db.Exec(queries.CreateCountry, country.Name, country.PeopleNumber)
	if err != nil {
		return -1, err
	} else {
		country.Id, _ = result.LastInsertId()
	}
	return country.Id, nil
}
func (database CountryDAOImpl) Update(country *Country) error {
	db := GetConnection()
	defer db.Close()
	_, err := db.Exec(queries.UpdateCountry, country.Name, country.PeopleNumber, country.Id)
	if err != nil {
		return err
	}
	return nil
}
func (database CountryDAOImpl) Delete(id int64) error {
	db := GetConnection()
	defer db.Close()
	_, err := db.Exec(queries.DeleteCountry, id)
	if err != nil {
		return err
	}
	return nil
}
