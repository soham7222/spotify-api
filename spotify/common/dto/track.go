package dto

type TrackDbModel struct {
	Isrc    string
	ImgURI  string
	Title   string
	Artists string
}

type TrackArtistsData struct {
	Artist []string `json:"artists"`
}
