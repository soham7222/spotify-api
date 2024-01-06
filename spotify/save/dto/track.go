package dto

type TrackDbModel struct {
	Isrc     string
	Metadata string
}

type TrackMetaData struct {
	ImgURI string   `json:"imageUri"`
	Title  string   `json:"title"`
	Artist []string `json:"artists"`
}
