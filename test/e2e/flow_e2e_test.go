package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

var baseURL = "http://localhost:8080/api/v1"

// For this test to run, the server and the database must be running
func TestFlowE2E(t *testing.T) {
	run := time.Now().Unix()

	primaryUser := map[string]string{
		"username": fmt.Sprintf("primary-%d", run),
		"email":    fmt.Sprintf("primary-%d@test.com", run),
		"password": "testpassword",
	}

	secondaryUser := map[string]string{
		"username": fmt.Sprintf("secondary-%d", run),
		"email":    fmt.Sprintf("secondary-%d@test.com", run),
		"password": "testpassword",
	}

	// Create primary user
	apiPrimaryUser := testCreateUser(t, primaryUser)
	apiPrimaryUserID := int64(apiPrimaryUser["id"].(float64))

	// Create secondary user
	apiSecondaryUser := testCreateUser(t, secondaryUser)
	apiSecondaryUserID := int64(apiSecondaryUser["id"].(float64))

	// Get token for primary user
	apiPrimaryUserLogin := testLogin(t, primaryUser)
	apiPrimaryUserToken := apiPrimaryUserLogin["token"].(string)

	// Get token for secondary user
	apiSecondaryUserLogin := testLogin(t, secondaryUser)
	apiSecondaryUserToken := apiSecondaryUserLogin["token"].(string)

	// Get primary user by ID
	testGetUserByID(t, apiPrimaryUserID, apiPrimaryUserToken)

	// Get secondary user by ID
	testGetUserByID(t, apiSecondaryUserID, apiSecondaryUserToken)

	// Create tweet for secondary user
	testCreateTweet(t, apiSecondaryUserID, apiSecondaryUserToken, run)

	// Get empty timeline for primary user
	testGetTimeline(t, apiPrimaryUserToken, true)

	// Create a follower from primary user to secondary user
	testCreateFollower(t, apiPrimaryUserID, apiSecondaryUserID, apiPrimaryUserToken)

	// Get not empty timeline for primary user
	testGetTimeline(t, apiPrimaryUserToken, false)
}

func testCreateUser(t *testing.T, user map[string]string) map[string]interface{} {
	jsonBody, _ := json.Marshal(user)

	res, err := http.Post(baseURL+"/user", "application/json", bytes.NewBuffer(jsonBody))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	assert.Equal(t, user["username"], response["username"])
	assert.NotNil(t, response["id"])

	return response
}

func testLogin(t *testing.T, user map[string]string) map[string]interface{} {
	body := map[string]string{
		"email":    user["email"],
		"password": user["password"],
	}
	jsonBody, _ := json.Marshal(body)

	res, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(jsonBody))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	assert.NotNil(t, response["token"])

	return response
}

func testGetUserByID(t *testing.T, userID int64, token string) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%d", baseURL, userID), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	assert.NotNil(t, response["id"])
	assert.Equal(t, userID, int64(response["id"].(float64)))
	assert.NotNil(t, response["username"])
	assert.Nil(t, response["password"])
}

func testCreateTweet(t *testing.T, userID int64, token string, run int64) {
	text := fmt.Sprintf("This is a test tweet from the user %d and run %d", userID, run)
	body := map[string]interface{}{
		"user_id": userID,
		"text":    text,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/tweet", baseURL), bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	assert.Equal(t, text, response["text"])
	assert.NotNil(t, response["id"])
	assert.NotNil(t, response["created_at"])

	return
}

func testGetTimeline(t *testing.T, token string, shouldBeEmpty bool) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/timeline", baseURL), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response []map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	if shouldBeEmpty {
		assert.Empty(t, response)
	} else {
		assert.NotEmpty(t, response)
	}
}

func testCreateFollower(t *testing.T, followerID int64, followedID int64, token string) {
	body := map[string]interface{}{
		"follower_id": followerID,
		"followed_id": followedID,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/follower", baseURL), bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := http.DefaultClient.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(res.Body).Decode(&response)

	assert.Equal(t, followerID, int64(response["follower_id"].(float64)))
	assert.Equal(t, followedID, int64(response["followed_id"].(float64)))
	assert.NotNil(t, response["created_at"])
}
