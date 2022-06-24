package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Record model
type Record struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	NID         string    `json:"nid"`
	CID         string    `json:"cid"`
	Metadata    string    `json:"-"`
	AccessToken string    `json:"accessToken"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// RecordMetadata struct
type RecordMetadata struct {
	Name         string    `json:"name" binding:"required"`
	Issuer       string    `json:"issuer"`
	IssuedAt     time.Time `json:"issuedAt"`
	SerialNo     string    `json:"serialNo"`
	ThumbnailURL string    `json:"thumbnailURL"`
	AssetURL     string    `json:"assetURL"`
}

// CreateRecordInput struct
type CreateRecordInput struct {
	NID      string         `json:"nid" binding:"required"`
	Metadata RecordMetadata `json:"metadata" binding:"required"`
}

// UpdateRecordInput struct
type UpdateRecordInput struct {
	NID string `json:"nid" binding:"required"`
}

// RecordOutput struct
type RecordOutput struct {
	ID          uint           `json:"id"`
	NID         string         `json:"nid"`
	NNS         string         `json:"nns"`
	AccessToken string         `json:"accessToken"`
	Metadata    RecordMetadata `json:"metadata"`
}

// GenerateRecordOutput method
func GenerateRecordOutput(record *Record, regNNS string, withAccessToken bool) RecordOutput {
	var metadata RecordMetadata
	json.Unmarshal([]byte(record.Metadata), &metadata)

	accessToken := "-"
	if withAccessToken {
		accessToken = record.AccessToken
	}

	return RecordOutput{
		ID:          record.ID,
		NID:         record.NID,
		NNS:         fmt.Sprintf("%d%s", record.ID, regNNS),
		Metadata:    metadata,
		AccessToken: accessToken,
	}
}
