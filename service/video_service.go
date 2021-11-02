package service

import (
	"github/quocdaitrn/golang-gin-poc/entity"
)

type VideoService interface {
	FindAll() []*entity.Video
	Save(*entity.Video)
}

type videoService struct {
	videos []*entity.Video
}

func New() VideoService {
	return &videoService{videos: []*entity.Video{}}
}

func (s *videoService) FindAll() []*entity.Video {
	return s.videos
}

func (s *videoService) Save(v *entity.Video) {
	s.videos = append(s.videos, v)
}
