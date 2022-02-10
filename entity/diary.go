package entity

import "time"

// Table represent books table in database
type Diary struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Body      string    `gorm:"type:text" json:"body"`
	CreatedAt time.Time `json:"createdat"`
	UserID    uint64    `gorm:"not null" json:"-"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
