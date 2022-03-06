package Teacher

import (
	"schoolmodel/Student"
)

type Teacher struct {
	FirstName     string
	LastName      string
	CourseName    string
	TeacherSalary float64
	GradeLevel    int32
	StudentList   map[int]Student.Student
	Result        []StudentResult
}

type StudentResult struct {
	firstName string
	lastName  string
	score     float64
}

func (t *Teacher) SetTeacherSalary(amount float64) {
	t.TeacherSalary = amount
}

//func (t *Teacher) UpdateStudentList(p *Principal.Principal) map[int]Student.Student {
//	for ind, val := range p.StudentList {
//		if val.Dept == t.CourseName {
//			t.StudentList[ind] = val
//		}
//	}
//	return t.StudentList
//}
