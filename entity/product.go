package entity

import (
	"database/sql"
	"time"
)

type Product struct {
	ID        int          `gorm:"primaryKey;column:product_id"`
	Name      string       `gorm:"unique;column:name;type:varchar(100)"`
	Price     float32      `gorm:"column:price"`
	Quantity  int32        `gorm:"column:quantity"`
	OwnerID   int          `gorm:"column:owner_id;not null"`
	CreatedAt time.Time    `gorm:"column:created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at"`
}

func (Product) TableName() string {
	return "tbl_product"
}
