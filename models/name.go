package models

import "gorm.io/gorm"

// Name model
type Name struct {
	gorm.Model
	NNS          uint   `gorm:"uniqueIndex" json:"nns"`
	NID          string `json:"nid" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Issuer       string `json:"issuer" binding:"required"`
	SerialNo     string `json:"serialNo"`
	ThumbnailURL string `json:"thumbnailURL"`
	AssetURL     string `json:"assetURL"`
	AccessToken  string `json:"accessToken"`
}
