package entities

type BeerStyle struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
	MinTemperature float64 `json:"min_temperature" gorm:"not null"`
	MaxTemperature float64 `json:"max_temperature" gorm:"not null"`
}