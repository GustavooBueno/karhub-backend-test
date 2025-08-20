package postgres

import (
	"github.com/GustavooBueno/karhub-backend-test/internal/domain/entities"
	"gorm.io/gorm"
)

type BeerStyleRepository struct {
	db *gorm.DB
}

func NewBeerStyleRepository(db *gorm.DB) *BeerStyleRepository {
	return &BeerStyleRepository{
		db: db,
	}
}

func (r *BeerStyleRepository) Create(style *entities.BeerStyle) error {
	return r.db.Create(style).Error
}

func (r *BeerStyleRepository) FindAll() ([]entities.BeerStyle, error) {
	var styles []entities.BeerStyle
	err := r.db.Find(&styles).Error
	return styles, err
}

func (r *BeerStyleRepository) FindByID(id uint) (*entities.BeerStyle, error) {
	var style entities.BeerStyle
	err := r.db.First(&style, id).Error
	return &style, err
}

func (r *BeerStyleRepository) Update(style *entities.BeerStyle) error {
	return r.db.Save(style).Error
}

func (r *BeerStyleRepository) Delete(id uint) error {
	return r.db.Delete(&entities.BeerStyle{}, id).Error
}