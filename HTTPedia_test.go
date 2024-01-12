package HTTPedia

import (
	"os"
	"reflect"
	"testing"
)

// TestParseStatusCodeFile tests the parseStatusCodeFile function.
func TestParseStatusCodeFile(t *testing.T) {
	// Setup - create a temporary test file
	testFilename := "test_status_codes.json"
	testContent := []byte(`[{"code":200,"description":"OK"},{"code":404,"description":"Not Found"}]`)
	os.WriteFile(testFilename, testContent, 0644)
	defer os.Remove(testFilename) // Cleanup after test

	// Execute
	statusCodes, err := parseStatusCodeFile(testFilename)

	// Verify
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(statusCodes) != 2 {
		t.Errorf("Expected 2 status codes, got %d", len(statusCodes))
	}

	expected := []StatusCode{
		{Code: 200, Description: "OK"},
		{Code: 404, Description: "Not Found"},
	}
	if !reflect.DeepEqual(statusCodes, expected) {
		t.Errorf("Expected %+v, got %+v", expected, statusCodes)
	}
}

// TestGetStatusCode tests the GetStatusCode function.
func TestGetStatusCode(t *testing.T) {
	cases := []struct {
		code     int
		expected *StatusCode
	}{
		{200, &StatusCode{200, "OK"}},
		{404, &StatusCode{404, "Not Found"}},
		{99, nil}, // 99 is not in our test file
	}

	for _, c := range cases {
		got, err := GetStatusCode(c.code)
		if c.expected == nil && err == nil {
			t.Errorf("Expected an error for code %d, got none", c.code)
		} else if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("GetStatusCode(%d): expected %+v, got %+v", c.code, c.expected, got)
		}
	}
}

