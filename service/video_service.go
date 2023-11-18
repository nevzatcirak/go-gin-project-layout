package service

import (
	"github.com/nevzatcirak/go-gin-poc/domain"
)

type VideoService interface {
	Save(domain.Video) error
	Update(domain.Video) error
	Delete(domain.Video) error
	FindAll() []domain.Video
}

type videoService struct {
	repository domain.VideoRepository
}

func NewVideoService(videoRepository domain.VideoRepository) VideoService {
	return &videoService{
		repository: videoRepository,
	}
}

func (service *videoService) Save(video domain.Video) error {
	service.repository.Save(video)
	return nil
}

func (service *videoService) Update(video domain.Video) error {
	service.repository.Update(video)
	return nil
}

func (service *videoService) Delete(video domain.Video) error {
	service.repository.Delete(video)
	return nil
}

func (service *videoService) FindAll() []domain.Video {
	return service.repository.FindAll()
}
