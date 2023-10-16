package services_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"time"
)

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "video_base.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}
