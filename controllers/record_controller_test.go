package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nidnetwork/nid-native-registry/controllers"
	"github.com/nidnetwork/nid-native-registry/helpers"
	"github.com/nidnetwork/nid-native-registry/models"
	"github.com/stretchr/testify/assert"
)

func TestRecord(t *testing.T) {
	models.ConnectDatabase(helpers.GetDSN())
	record := createRecordTest(t)
	getRecordTest(t, record)
	updateOutput := updateRecordTest(t, record)
	deleteRecordTest(t, updateOutput)
}

func createRecordTest(t *testing.T) (record models.Record) {
	router := controllers.CreateRouter()
	w := httptest.NewRecorder()
	nid := "did:nid:eip155_9:erc721_3:9"
	body, _ := json.Marshal(models.CreateRecordInput{
		NID: nid,
		Metadata: models.RecordMetadata{
			Name: "My awesome NFT",
		},
	})
	reader := strings.NewReader(string(body))
	req, _ := http.NewRequest("POST", "/api/v1/records", reader)
	router.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &record); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, nid, record.NID)

	return record
}

func getRecordTest(t *testing.T, record models.Record) {
	router := controllers.CreateRouter()
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/api/v1/records/%d", record.ID)
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	var output models.RecordOutput
	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, record.NID, output.NID)
}

func updateRecordTest(t *testing.T, record models.Record) (output models.Record) {
	router := controllers.CreateRouter()
	w := httptest.NewRecorder()
	nid := "did:nid:eip155_9:erc721_3:10"
	body, _ := json.Marshal(models.UpdateRecordInput{
		NID: nid,
	})
	reader := strings.NewReader(string(body))
	url := fmt.Sprintf("/api/v1/records/%d", record.ID)
	req, _ := http.NewRequest("PATCH", url, reader)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", record.AccessToken))
	router.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, nid, output.NID)

	return output
}

func deleteRecordTest(t *testing.T, record models.Record) {
	router := controllers.CreateRouter()
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/api/v1/records/%d", record.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", record.AccessToken))
	router.ServeHTTP(w, req)

	var result map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "yes", result["success"])
}
