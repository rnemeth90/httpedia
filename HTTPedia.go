package HTTPedia

import (
	"encoding/json"
	"fmt"
	"os"
)

type StatusCode struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func GetStatusCode(code int) (*StatusCode, error) {
	statusCodesFromFile, err := parseStatusCodeFile("status_codes.json")
	if err != nil {
		return nil, err
	}

	for _, statusCode := range statusCodesFromFile {
		if statusCode.Code == code {
			return &statusCode, nil
		}
	}

	return nil, fmt.Errorf("status code not found: %d", code)
}

func parseStatusCodeFile(filename string) ([]StatusCode, error) {
	statusCodes := []StatusCode{}

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(file, &statusCodes); err != nil {
		return nil, err
	}

	return statusCodes, nil
}
