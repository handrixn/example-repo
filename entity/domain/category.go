package domain

import (
	"github.com/google/uuid"
	"github.com/handrixn/example-repo/helper"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID string `json:"uuid" gorm:"size:36;not null"`
	Name string `json:"name" gorm:"unique;size:255;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func (c *Category) TableName() string {
	return "categories"
}

func (c *Category) BeforeCreate(db *gorm.DB) error {
	newUUID, err := uuid.NewRandom()
	helper.PanicIfError(err)

	c.UUID = newUUID.String()

	return err
}