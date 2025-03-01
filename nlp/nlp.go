package nlp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GeminiResponse represents the structure of the Google Gemini API response.
type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// ParseInstruction sends the instruction to the Google Gemini API and returns a shell command.
func ParseInstruction(apiKey, model, instruction string) (string, error) {
	// Prompt instructs the model to output only the shell command.
	prompt := fmt.Sprintf(
		"Translate the following natural language command into a safe shell command. "+
			"Only output the shell command and nothing else.\n\n%s", instruction)

	// Construct the request body for Gemini API.
	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	// Construct the Gemini API URL with the model and API key.
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, apiKey)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Gemini API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the response.
	var geminiResponse GeminiResponse
	if err := json.Unmarshal(body, &geminiResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal API response: %v", err)
	}

	if len(geminiResponse.Candidates) == 0 {
		return "", fmt.Errorf("no candidates in response")
	}
	if len(geminiResponse.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no parts in content")
	}

	// Extract the generated shell command.
	shellCommand := geminiResponse.Candidates[0].Content.Parts[0].Text
	return shellCommand, nil
}
