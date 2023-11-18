package domain

import "time"

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" bson:"title" binding:"min=2,max=100" gorm:"type:varchar(100)" validate:"is-cool"`
	Description string    `json:"description" bson:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" bson:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      User      `json:"author" bson:"author" gorm:"foreignkey:UserID"`
	UserID      uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at" gorm:"default:CURRENT_TIMESTAMP" `
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type VideoRepository interface {
	Save(video Video)
	Update(video Video)
	Delete(video Video)
	FindAll() []Video
}
