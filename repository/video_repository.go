package repository

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nevzatcirak/go-gin-poc/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type videoRepository struct {
	connection *gorm.DB
}

func NewVideoRepository(db *gorm.DB) domain.VideoRepository {
	return &videoRepository{
		connection: db,
	}
}

func (db *videoRepository) Save(video domain.Video) {
	db.connection.Omit(clause.Associations).Create(&video)
}

func (db *videoRepository) Update(video domain.Video) {
	db.connection.Omit(clause.Associations).Save(&video)
}

func (db *videoRepository) Delete(video domain.Video) {
	db.connection.Delete(&video)
}

func (db *videoRepository) FindAll() []domain.Video {
	var videos []domain.Video
	db.connection.Preload(clause.Associations).Find(&videos)
	return videos
}
