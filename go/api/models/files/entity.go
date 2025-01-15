package files

import (
	"time"

	"gorm.io/gorm"
)

type FileEntity struct {
	gorm.Model
	Id          uint      `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Filename    string    `json:"filename" db:"filename"`
	Path        string    `json:"-" db:"path"`
	Size        int64     `json:"size" db:"size"`
	MimeType    string    `json:"mime_type" db:"mime_type"`
	Description string    `json:"description" db:"description"`
	Tags        string    `json:"tags" db:"tags"`
	Url         string    `json:"url" db:"url"`
	UploadedAt  time.Time `json:"uploaded_at" db:"uploaded_at"`
}
