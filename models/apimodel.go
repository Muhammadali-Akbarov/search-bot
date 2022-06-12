package models

type SimpleStruct struct {
	Value []struct {
		ThumbnailUrl string `json:"thumbnailUrl"`
	}
}
