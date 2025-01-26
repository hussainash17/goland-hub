package models

import (
	"time"
)

type News struct {
	NewsID    uint      `gorm:"column:news_id;primaryKey" json:"news_id"` // Maps to "news_id" in DB
	FirmID    string    `gorm:"column:firm_id" json:"firm_id"`            // Maps to "firm_id" in DB
	Title     string    `gorm:"column:title" json:"title"`                // Maps to "title" in DB
	Reference string    `gorm:"column:reference" json:"reference"`        // Maps to "reference" in DB
	NewsText  string    `gorm:"column:news_text" json:"news_text"`        // Maps to "news_text" in DB
	Timestamp time.Time `gorm:"column:timestamp" json:"timestamp"`        // Maps to "timestamp" in DB
	StockCode string    `gorm:"column:stock_code" json:"stock_code"`      // Maps to "stock_code" in DB
}

// TableName maps the struct to the table name in the database
func (News) TableName() string {
	return "news"
}
