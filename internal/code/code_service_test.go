package code_test

import (
	"testing"

	"github.com/eduardylopes/link-shortener/internal/code"
)

func TestGenerateUniqueCode(t *testing.T) {
	s := code.NewService()

	t.Run("TestLength", func(t *testing.T) {
		code := s.GenerateUniqueCode()

		if len(code) != 6 {
			t.Errorf("Expected code length 6, but got %d", len(code))
		}
	})

	t.Run("TestUniqueness", func(t *testing.T) {
		uniqueCodes := make(map[string]bool)

		for i := 0; i < 1000; i++ {
			code := s.GenerateUniqueCode()

			if uniqueCodes[code] {
				t.Errorf("Duplicate code generated: %s", code)
			}

			uniqueCodes[code] = true
		}
	})

}
