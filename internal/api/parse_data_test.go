package api

import (
	"testing"
)

func TestParseData(t *testing.T) {
	validJSON := []byte(`{
		"count": 1,
		"results": [{"name": "area", "url": "http://example.com/area"}]
	}`)

	invalidJSON := []byte(`{invalid}`)

	t.Run("Valid JSON", func(t *testing.T) {
		areas := LocationAreas{}

		result, err := ParseData(validJSON, &areas)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result.Count != 1 {
			t.Errorf("Expected count 1, got %d", result.Count)
		}

		if len(result.Results) != 1 || result.Results[0].Name != "area" {
			t.Errorf("Unexpected results: %+v", result.Results)
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		areas := LocationAreas{}

		_, err := ParseData(invalidJSON, &areas)

		if err == nil {
			t.Fatal("expected an error but got none")
		}
	})
}
