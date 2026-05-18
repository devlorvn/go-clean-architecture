package model

type Post struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Title    string `gorm:"type:varchar(255);not null"`
	Content  string `gorm:"type:text;not null"`
	AuthorID int64  `gorm:"not null"`
}
