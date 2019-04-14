package interfaces

import "../entities"

type CountryDAO interface {
	Create(country *entities.Country) (int64, error)
	Read(id int64) (*entities.Country, error)
	ReadAll() ([]*entities.Country, error)
	Update(country *entities.Country) error
	Delete(id int64) error
}
