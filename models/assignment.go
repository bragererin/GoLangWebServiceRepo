package models

type Assignment struct {
	StudentID int
	TestScore int
}

var (
	assignment []*Assignment
)

func AddAssignment(a Assignment) (Assignment, error) {	
	assignment = append(assignment, &a)
	return a, nil
}

func GetAverageComputation(a []Assignment) float64 {
	var total float64 = 0
	for i, v := range a {
		total += float64(v.TestScore)
		i++
	}

	return total / float64(len(a))
}
