package model

type TrackDetailsResponse struct {
	Isrc    string   `json:"isrc"`
	ImgURI  string   `json:"imageUri"`
	Title   string   `json:"title"`
	Artists []string `json:"artists"`
}
