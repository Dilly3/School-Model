package Student

import (
	"schoolmodel/Course"
)

type Student struct {
	FirstName  string
	LastName   string
	SchoolFees float64
	Matno      int32
	CourseList []Course.Course
}

func (s *Student) AddCourse(course Course.Course) {

}
