package School

import (
	"fmt"
	"schoolmodel/Department"
	"schoolmodel/Student"
	"schoolmodel/Teacher"
	"strings"
)

type amount = float64
type deptName = string

//********************************************
type School struct {
	TeachersList map[int]Teacher.Teacher
	StudentsList map[int]Student.Student
	Departs      []Department.Department
}

func (sch *School) getTeacherList() map[int]Teacher.Teacher {
	newMap := make(map[int]Teacher.Teacher)
	for ind, val := range sch.TeachersList {
		fmt.Println(val)
		newMap[ind] = val
	}
	return newMap
}
func (sch *School) EmployTeacher(teacher *Teacher.Teacher) {
	teachers := len(sch.TeachersList)
	sch.TeachersList[teachers+1] = *teacher
}

func (sch *School) SackTeacher(teacher Teacher.Teacher) {
	name1 := teacher.FirstName + " " + teacher.LastName
	name1 = strings.ToLower(name1)
	for ind, val := range sch.TeachersList {
		name2 := val.FirstName + " " + val.LastName
		name2 = strings.ToLower(name2)
		if name1 == name2 {
			delete(sch.TeachersList, ind)
		}
	}
}
func (sch *School) PayTeacherSalary(teacher *Teacher.Teacher, amount float64) map[int]Teacher.Teacher {

	var temp Teacher.Teacher
	for ind, val := range sch.getTeacherList() {

		if (val.FirstName == teacher.FirstName) && (val.LastName == teacher.LastName) {
			temp = Teacher.Teacher{val.FirstName, val.LastName, val.CourseName,
				amount, val.GradeLevel, val.StudentList, val.Result}
			sch.TeachersList[ind] = temp
		}

	}

	return sch.TeachersList
}
