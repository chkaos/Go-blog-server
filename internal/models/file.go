package models

type File struct {
	Model

	FileName string `json:"file_name"`
	URL      string `json:url`
	Type     string `json:"type"`
	Size     int    `json:"size"`
}


