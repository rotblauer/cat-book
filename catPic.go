package main

import (
	"github.com/rotblauer/trackpoints/trackPoint"
	"image"
)

type CatPic struct {
	TrackPoint trackPoint.TrackPoint `json:"trackpoint"` //the cat track
	CatPic     image.Image           `json:"catpic"`
}
