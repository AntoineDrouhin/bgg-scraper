package scraper

import (
	"testing"
)

func TestExtractGameName(t *testing.T) {

	t.Run("Great Western Trail (2016)", func(t *testing.T) {
		expectedName := "Great Western Trail"
		expectedDate := 2016
		name, date := ExtractGameName("Great Western Trail (2016)")

		if name != expectedName {
			t.Errorf("Extracted name incorrect, got: %s want %s", name, expectedName)
		}
		if date != expectedDate {
			t.Errorf("Extracted date incorrect, got: %d want %d", date, expectedDate)

		}
	})

	t.Run("Twilight Imperium (Fourth Edition) (2017)", func(t *testing.T) {
		expectedName := "Twilight Imperium (Fourth Edition)"
		expectedDate := 2017
		name, date := ExtractGameName("Twilight Imperium (Fourth Edition) (2017)")

		if name != expectedName {
			t.Errorf("Extracted name incorrect, got: %s want %s", name, expectedName)
		}

		if date != expectedDate {
			t.Errorf("Extracted date incorrect, got: %d want %d", date, expectedDate)

		}
	})

}
