package main

import (
	"fmt"
	"schoolmodel/Applicant"
	"schoolmodel/Department"
	"schoolmodel/Principal"
	"schoolmodel/School"
	"schoolmodel/Student"
	"schoolmodel/Teacher"
)

type amount = float64

func main() {
	departs1 := Department.GetDepartments()
	//var stud1 = Student.Student{"john", "obi", "medicine", 0,
	//	0, map[int]Course.Course{}, false}
	//var stud2 = Student.Student{"chidi", "lawal", "engineering", 0,
	//	0, map[int]Course.Course{}, false}

	var apllicant1 = Applicant.Applicant{"chinedu", "ofudu", 250, "fine arts"}
	var apllicant2 = Applicant.Applicant{"james", "king", 295, "medicine"}

	var principal1 = Principal.Principal{map[int]Student.Student{}, map[int]Department.Department{}}

	var mySch = School.School{
		map[int]Teacher.Teacher{}, map[int]Student.Student{}, departs1,
	}
	mysch := &mySch
	var teacher1 = Teacher.Teacher{"francis", "jide", "medicine",
		0, 0, map[int]Student.Student{}, []Teacher.StudentResult{}}
	var teacher2 = Teacher.Teacher{"ngozi", "okonjo", "engineering",
		0, 0, map[int]Student.Student{}, []Teacher.StudentResult{}}

	mySch.EmployTeacher(&teacher2)
	mySch.EmployTeacher(&teacher1)

	mySch.PayTeacherSalary(&teacher1, 50000)
	mySch.PayTeacherSalary(&teacher2, 800000)
	mySch.EmployTeacher(&teacher2)
	mySch.EmployTeacher(&teacher1)

	mySch.PayTeacherSalary(&teacher1, 50000)
	mySch.PayTeacherSalary(&teacher2, 800000)
	fmt.Println("*************************************************\n")
	///************************************************************

	principal1.AdmitStudent(mysch, apllicant1)
	principal1.AdmitStudent(mysch, apllicant2)
	principal1.ExpelStudent(mysch, "chinedu ofudu")

	fmt.Println(mysch.StudentsList)
	principal1.UpdateStudentList(mysch)
	principal1.AdmitStudent(mysch, apllicant1)
	principal1.UpdateStudentList(mysch)
	fmt.Println(principal1.StudentList)

}
