package models_test

import (
	"Itential/AverageComputation/GoLangWebService/models"
	"testing"
)

// TestCase 1: Add test scores for a class
func GetAverageComputationTest(t *testing.T) {
	mathTestScoresForClassPeriodA := []models.Assignment{
		{
			StudentID: 123456,
			TestScore: 100,
		},
		{
			StudentID: 234567,
			TestScore: 60,
		},
		{
			StudentID: 345678,
			TestScore: 72,
		},
		{
			StudentID: 456789,
			TestScore: 83,
		},
		{
			StudentID: 567890,
			TestScore: 95,
		},
		{
			StudentID: 678901,
			TestScore: 87,
		},
	}
	got := models.GetAverageComputation(mathTestScoresForClassPeriodA)
	expect := 82.83

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %f, got: %f\n", expect, got)
	}
}
