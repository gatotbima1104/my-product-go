package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductsType string

const (
	ProductsTypeScraping   = "scraping"
	ProductsTypeAutomation = "automation"
)

type Products struct {
	ID          uuid.UUID
	Name        string
	Price       float64
	Type        ProductsType
	Description string
	Photo       string
	Created_at  time.Time
	Updated_at  time.Time
}