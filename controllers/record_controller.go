package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidnetwork/nid-native-registry/helpers"
	"github.com/nidnetwork/nid-native-registry/models"
	"github.com/nidnetwork/nid-native-registry/utils"
)

// FindRecords godoc
// @Summary Find all records
// @Schemes
// @Description Find current account all records
// @Tags Record
// @Accept json
// @Produce json
// @Success 200 {array} models.Record
// @Router /records [get]
func FindRecords(c *gin.Context) {
	var records []models.Record
	models.DB.Find(&records)

	c.JSON(http.StatusOK, records)
}

// GetRecord godoc
// @Summary Get a record
// @Schemes
// @Description Get a record
// @Tags Record
// @Accept json
// @Produce json
// @Param id path int true  "Record ID"
// @Success 200 {object} models.RecordOutput
// @Router /records/{id} [get]
func GetRecord(c *gin.Context) {
	record, err := helpers.GetRecordByID(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, models.GenerateRecordOutput(record))
}

// CreateRecord godoc
// @Summary Create new record
// @Schemes
// @Description Create new record
// @Tags Record
// @Accept json
// @Produce json
// @Param record body models.CreateRecordInput true "Create new record"
// @Success 200 {object} models.Record
// @Router /records [post]
func CreateRecord(c *gin.Context) {
	// Validate input
	var input models.CreateRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.NewBadRequestError(c, err, "BindJSONError")
		return
	}

	// Create record
	record := models.Record{
		NID:          input.NID,
		Name:         input.Name,
		Issuer:       input.Issuer,
		IssuedAt:     input.IssuedAt,
		SerialNo:     input.SerialNo,
		ThumbnailURL: input.ThumbnailURL,
		AssetURL:     input.AssetURL,
		AccessToken:  utils.RandomString(30),
	}

	models.DB.Create(&record)

	c.JSON(http.StatusOK, record)
}

// UpdateRecord godoc
// @Summary Update a record
// @Schemes
// @Description Update a record
// @Tags        Record
// @Accept      json
// @Produce     json
// @Param       id       path     int                 true  "Record ID"
// @Param       record body     models.UpdateRecordInput true "Update a record"
// @Success     200      {object} models.Record
// @Security    BearerAuth
// @Router      /records/{id} [patch]
func UpdateRecord(c *gin.Context) {
	record, err := helpers.GetRecordByID(c)
	if err != nil {
		return
	}

	err = helpers.RecordAuth(c, record)
	if err != nil {
		return
	}

	var input models.UpdateRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.NewBadRequestError(c, err, "BindJSONError")
		return
	}

	models.DB.Model(&record).Updates(models.Record{
		NID:         input.NID,
		AccessToken: utils.RandomString(30),
	})

	c.JSON(http.StatusOK, record)
}

// DeleteRecord godoc
// @Summary     Delete a record
// @Schemes
// @Description Delete a record
// @Tags        Record
// @Accept      json
// @Produce     json
// @Param       id path int true  "Record ID"
// @Success     204  {object}  models.Record
// @Security    BearerAuth
// @Router      /records/{id} [delete]
func DeleteRecord(c *gin.Context) {
	record, err := helpers.GetRecordByID(c)
	if err != nil {
		return
	}

	err = helpers.RecordAuth(c, record)
	if err != nil {
		return
	}

	models.DB.Delete(&record)

	c.JSON(http.StatusOK, gin.H{"success": "yes"})
}