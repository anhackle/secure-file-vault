package dto

type FileID struct {
	FileID string `json:"file_id" binding:"required,uuid4"`
}

type ShareUri struct {
	ID string `uri:"id" binding:"required,uuid"`
}
