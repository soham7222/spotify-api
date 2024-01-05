package model

type SaveSongRequest struct {
	ISRC string `json:"isrc"`
}

type TracksSearchResponse struct {
	Tracks TracksCollection `json:"tracks"`
}

type TracksCollection struct {
	Items []TracksDetailsItem `json:"items"`
	Total int                 `json:"total"`
}

type TracksDetailsItem struct {
	Album       TrackAlbumDetails    `json:"album"`
	Artists     []TrackArtistDetails `json:"artists"`
	ExternalIds TrackExternalIds     `json:"external_ids"`
	Name        string               `json:"name"`
	Popularity  int                  `json:"popularity"`
}

type TrackAlbumDetails struct {
	Images []TrackAlbumImageDetails `json:"images"`
}

type TrackAlbumImageDetails struct {
	URL string `json:"url"`
}

type TrackArtistDetails struct {
	Name string `json:"name"`
}

type TrackExternalIds struct {
	Isrc string `json:"isrc"`
}
