package model

import "time"

type Post struct {
	ID      int64  `gorm:"primaryKey;autoIncrement"`
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:text;not null"`

	AuthorID  int64     `gorm:"not null"`
	Author    Author    `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
