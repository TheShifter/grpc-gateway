package queries

const (
	CreateCountry  = `INSERT INTO country_schema.country(name, peopleNumber) VALUES(?, ?)`
	GetCountryById = `SELECT * FROM country_schema.country WHERE id = ?`
	GetCountries   = `SELECT * FROM country_schema.country`
	UpdateCountry  = `UPDATE country_schema.country SET name = ?, peopleNumber = ?  WHERE id = ?`
	DeleteCountry  = `DELETE FROM country_schema.country WHERE id = ?`
)
