package files

import "time"

type FileEntity struct {
	ID          string    `json:"id" db:"id"`
	Filename    string    `json:"filename" db:"filename"`
	Path        string    `json:"-" db:"path"`
	Size        int64     `json:"size" db:"size"`
	MimeType    string    `json:"mime_type" db:"mime_type"`
	Description string    `json:"description" db:"description"`
	Tags        string    `json:"tags" db:"tags"`
	UploadedAt  time.Time `json:"uploaded_at" db:"uploaded_at"`
}
