package files

import "time"

type FileResponse struct {
	ID          string    `json:"id"`
	Filename    string    `json:"filename"`
	Size        int64     `json:"size"`
	MimeType    string    `json:"mime_type"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"`
	UploadedAt  time.Time `json:"uploaded_at"`
}

type UploadResponse struct {
	Success  bool         `json:"success"`
	Message  string       `json:"message"`
	FileInfo FileResponse `json:"file_info"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    string `json:"code,omitempty"`
}
