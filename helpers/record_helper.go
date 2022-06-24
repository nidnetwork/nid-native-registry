package helpers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nidnetwork/nid-native-registry/models"
)

// GetRecordByID that find name by ID
func GetRecordByID(c *gin.Context) (*models.Record, error) {
	var name models.Record
	if err := models.DB.Where("id = ?", c.Param("id")).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return nil, err
	}

	return &name, nil
}

// GetRecordByNNS that find name by NNS
func GetRecordByNNS(c *gin.Context) (*models.Record, error) {
	id := strings.Trim(c.Param("nns"), GetRegistarNNS())
	var name models.Record
	if err := models.DB.Where("id = ?", id).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return nil, err
	}

	return &name, nil
}

// RecordAuth auth required
func RecordAuth(c *gin.Context, name *models.Record) error {
	token := c.GetHeader("Authorization")
	if token == "" {
		err := errors.New("Empty authorization header")
		NewUnauthorizedError(c, err, "EmptyAuthorizationHeader")
		return err
	}

	accessToken := strings.Split(token, "Bearer ")
	if len(accessToken) < 2 {
		err := errors.New("Must provide Authorization header with format `Bearer {token}`")
		NewUnauthorizedError(c, err, "AuthorizationHeaderFormatError")
		return err
	}

	if accessToken[1] != name.AccessToken {
		err := errors.New("Invalid access token")
		NewUnauthorizedError(c, err, "InvalidAccessToken")
		return err
	}

	return nil
}
