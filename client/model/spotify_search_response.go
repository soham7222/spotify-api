package model

import (
	"encoding/json"
	"spotify-api/dto"
)

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

func (res TracksSearchResponse) TransformToDbModel(isrc string) dto.TrackDbModel {
	mostPopularTrack := res.getTheMostPopularTrack()
	artistData := dto.TrackArtistsData{
		Artist: mostPopularTrack.getArtists(),
	}

	data, _ := json.Marshal(artistData)

	return dto.TrackDbModel{
		Isrc:    isrc,
		Artists: string(data),
		ImgURI:  mostPopularTrack.getImageUri(),
		Title:   mostPopularTrack.Name,
	}
}

func (res TracksSearchResponse) getTheMostPopularTrack() TracksDetailsItem {
	if len(res.Tracks.Items) == 0 {
		return TracksDetailsItem{}
	}

	highest := res.Tracks.Items[0]

	if len(res.Tracks.Items) == 1 {
		return highest
	}

	for _, item := range res.Tracks.Items {
		if item.Popularity > highest.Popularity {
			highest = item
		}
	}
	return highest
}

func (track TracksDetailsItem) getImageUri() string {
	if len(track.Album.Images) > 0 {
		return track.Album.Images[0].URL
	}

	return ""
}

func (track TracksDetailsItem) getArtists() []string {
	result := make([]string, 0)

	if len(track.Artists) > 0 {
		for _, item := range track.Artists {
			result = append(result, item.Name)
		}

		return result
	}

	return nil
}
