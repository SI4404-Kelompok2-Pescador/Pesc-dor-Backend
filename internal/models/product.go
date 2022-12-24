package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID      string      `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name    string      `json:"name"`
	Price   float64     `json:"price"`
	StoreID string      `json:"store_id"`
	Store   Store `json:"store" gorm:"foreignKey:StoreID"`

}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}