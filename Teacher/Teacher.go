package Teacher

import "schoolmodel/Student"

type Teacher struct {
	TeacherName   string
	CourseName    string
	TeacherSalary float64
	GradeLevel    int32
	StudentList   []Student.Student
	Result        []result
}

type result struct {
	firstName string
	lastName  string
	score     float64
}
