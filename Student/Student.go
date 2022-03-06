package Student

import (
	"schoolmodel/Course"
	"schoolmodel/Department"
)

type Student struct {
	FirstName      string
	LastName       string
	Dept           string
	SchoolFees     float64
	Matno          int
	CourseList     map[int]Course.Course
	PaidFeesInFull bool
}

func (s *Student) UpdateFeesStatus() {
	// check student school fees paid
	// check student dept fees
	// compare fees paid with dept fees
	schfees := Department.GetDepartmentFees(s.Dept)
	if s.SchoolFees >= schfees {
		s.PaidFeesInFull = true
	}

}
