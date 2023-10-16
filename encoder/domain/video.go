package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Video struct {
	ID         string    `valid:"uuid" gorm:"type:uuid;primary_key" json:"enconded_video_folder"`
	ResourceID string    `valid:"-" gorm:"type:varchar(255)" json:"resource_id"`
	FilePath   string    `valid:"notnull" gorm:"type:varchar(255)" json:"file_path"`
	CreatedAt  time.Time `valid:"-" json:"-"`
	Jobs       []*Job    `valid:"-" gorm:"ForeignKey:VideoID" json:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}

	return nil
}
