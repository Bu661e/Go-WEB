package models

import (
	"testing"
)

func TestStudentType(t *testing.T) {

	s := Student{
		ID:   1,
		Name: "Bu661e",
		Age:  18,
	}
	stuType := s.GetType()

	if stuType != "famale" {
		t.Errorf("Expected 'famale', got '%s'", stuType)
	}

}
