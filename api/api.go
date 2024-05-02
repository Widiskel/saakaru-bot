package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	model "saakaru/model"
	"time"
)

func ExtractData(resp *http.Response) ([]byte, error) {
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Close the response body
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response body into a map
	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return nil, err
	}

	code, ok := responseBody["code"].(float64) // Assuming "code" is a number
	if !ok {
		return nil, fmt.Errorf("code field not found in API response")
	}

	// Check if the code is not zero
	if code != 0 {
		return nil, fmt.Errorf("non-zero code in API response: %d", int(code))
	}

	// Extract the "data" field from the response
	dataJSON, ok := responseBody["data"]
	if !ok {
		return nil, fmt.Errorf("data field not found in API response")
	}

	// Marshal the "data" field back into JSON
	dataBody, err := json.Marshal(dataJSON)
	if err != nil {
		return nil, err
	}

	return dataBody, nil
}

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func GetQuest(token string) (*model.QuestModel, error) {
	url := "https://api-saakuru-gainz.beyondblitz.app/blitz/quest/current-quest"
	requestBody := map[string]interface{}{
		"withQuestTask": true,
		"withStatistic": true,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", "Bearer "+token)
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dataBody, err := ExtractData(resp)
	if err != nil {
		return nil, err
	}

	var questResponse model.QuestModel
	err = json.Unmarshal(dataBody, &questResponse)
	if err != nil {
		return nil, err
	}

	return &questResponse, nil
}

func GetUser(token string) (*model.UserModel, error) {
	url := "https://api-saakuru-gainz.beyondblitz.app/blitz/user/current-profile"

	jsonBody, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", "Bearer "+token)
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dataBody, err := ExtractData(resp)
	if err != nil {
		return nil, err
	}

	var userResponse model.UserModel
	err = json.Unmarshal(dataBody, &userResponse)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func GetUserStatistic(token string, questId string) (*model.UserStatisticModel, error) {
	url := "https://api-saakuru-gainz.beyondblitz.app/blitz/quest/get-user-quest-statistic"
	requestBody := map[string]interface{}{
		"questId": questId,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", "Bearer "+token)
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dataBody, err := ExtractData(resp)
	if err != nil {
		return nil, err
	}

	var questResponse model.UserStatisticModel
	err = json.Unmarshal(dataBody, &questResponse)
	if err != nil {
		return nil, err
	}

	return &questResponse, nil
}

func ClaimQuest(token string, questId string, taskId string) (*string, error) {
	url := "https://api-saakuru-gainz.beyondblitz.app/blitz/quest/claim-quest-task"
	requestBody := map[string]interface{}{
		"questId": questId,
		"taskIds": taskId,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("authorization", "Bearer "+token)
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("=> Task " + taskId + " Claimed Successfully")
		success := "success"
		return &success, nil
	} else {
		fmt.Println("=> Failed to Complete & Claim Task " + taskId)
		failed := "failed"
		return &failed, nil
	}
}
