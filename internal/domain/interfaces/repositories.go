package interfaces

import "github.com/GustavooBueno/karhub-backend-test/internal/domain/entities"

type BeerStyleRepository interface {
    Create(style *entities.BeerStyle) error
    FindAll() ([]entities.BeerStyle, error)
    FindByID(id uint) (*entities.BeerStyle, error)
    Update(style *entities.BeerStyle) error
    Delete(id uint) error
}