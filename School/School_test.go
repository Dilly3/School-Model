package School

import (
	"schoolmodel/Department"
	"schoolmodel/Student"
	"schoolmodel/Teacher"
	"testing"
)

var departs1 = Department.GetDepartments()
var mysch = &mySch
var mySch = School{

	map[int]Teacher.Teacher{}, map[int]Student.Student{}, departs1,
}

func TestSchool_EmployTeacher(t *testing.T) {
	table := []struct {
		input     Teacher.Teacher
		expected1 string
		expected2 string
	}{
		{Teacher.Teacher{"john", "olisa", "medicine",
			0, 8, map[int]Student.Student{}}, "john", "olisa"},
		{Teacher.Teacher{"james", "oluwa", "medicine",
			0, 8, map[int]Student.Student{}}, "james", "oluwa"},
	}

	for _, val := range table {
		mySch.EmployTeacher(&val.input)

		if !(mySch.TeachersList[len(mySch.StudentsList)-1].FirstName == val.expected1) && (mySch.TeachersList[len(mySch.StudentsList)-1].LastName == val.expected2) {
			t.Errorf("Expected : %s %s found %s %s",
				val.expected1, val.expected2, mySch.TeachersList[len(mySch.StudentsList)-1].FirstName, mySch.TeachersList[len(mySch.StudentsList)-1].LastName)
		}
	}
}

var mySch2 = School{

	map[int]Teacher.Teacher{}, map[int]Student.Student{}, departs1,
}

//func TestSchool_PayTeacherSalary(t *testing.T) {
//	table := []struct {
//		input    Teacher.Teacher
//		expected float64
//	}{
//		{Teacher.Teacher{"john", "olisa", "medicine",
//			0, 8, map[int]Student.Student{}}, 50000},
//	}
//
//	for _, val := range table {
//		mySch2.EmployTeacher(&val.input)
//		list := mySch2.getTeacherList()
//		mySch2.PayTeacherSalary(&val.input, val.expected)
//
//		if list[len(list)-1].TeacherSalary != val.expected {
//			t.Errorf("Expected %v : got %v", val.expected, mySch2.TeachersList[len(list)-1].TeacherSalary)
//		}
//	}
//}
