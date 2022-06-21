package models

import (
	"time"
)

// Record model
type Record struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	NID          string    `json:"nid"`
	Name         string    `json:"name"`
	Issuer       string    `json:"issuer"`
	IssuedAt     time.Time `json:"issuedAt"`
	SerialNo     string    `json:"serialNo"`
	ThumbnailURL string    `json:"thumbnailURL"`
	AssetURL     string    `json:"assetURL"`
	AccessToken  string    `json:"accessToken"`
}

// CreateRecordInput struct
type CreateRecordInput struct {
	NID          string    `json:"nid" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Issuer       string    `json:"issuer"`
	IssuedAt     time.Time `json:"issuedAt"`
	SerialNo     string    `json:"serialNo"`
	ThumbnailURL string    `json:"thumbnailURL"`
	AssetURL     string    `json:"assetURL"`
}

// UpdateRecordInput struct
type UpdateRecordInput struct {
	NID string `json:"nid" binding:"required"`
}

// RecordOutput struct
type RecordOutput struct {
	ID           uint      `json:"id"`
	NID          string    `json:"nid"`
	Name         string    `json:"name"`
	Issuer       string    `json:"issuer"`
	IssuedAt     time.Time `json:"issuedAt"`
	SerialNo     string    `json:"serialNo"`
	ThumbnailURL string    `json:"thumbnailURL"`
	AssetURL     string    `json:"assetURL"`
}

// GenerateRecordOutput method
func GenerateRecordOutput(record *Record) RecordOutput {
	return RecordOutput{
		ID:           record.ID,
		NID:          record.NID,
		Name:         record.Name,
		Issuer:       record.Issuer,
		IssuedAt:     record.IssuedAt,
		SerialNo:     record.SerialNo,
		ThumbnailURL: record.ThumbnailURL,
		AssetURL:     record.AssetURL,
	}
}
