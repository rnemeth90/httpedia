package HTTPedia

import (
	"encoding/json"
	"os"
)

type StatusCode struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func GetStatusCode(code int) (string, error) {
	statusCodesFromFile, err := parseJson("status_codes.json")
	if err != nil {
		return "", err
	}

	returnValue := StatusCode{}

	if len(statusCodesFromFile) > 0 {
		for _, statusCode := range statusCodesFromFile {
			if statusCode.Code == code {
				returnValue = statusCode
			}
		}
	}

	return returnValue.Description, nil
}

func parseJson(filename string) ([]StatusCode, error) {
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
