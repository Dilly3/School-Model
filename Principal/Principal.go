package Principal

import (
	"fmt"
	"schoolmodel/Applicant"
	"schoolmodel/Course"
	"schoolmodel/Department"
	"schoolmodel/School"
	"schoolmodel/Student"
)

type Principal struct {
	StudentList map[int]Student.Student
	Depts       map[int]Department.Department
}

func (p *Principal) AdmitStudent(sch *School.School, jambite Applicant.Applicant) {
	depts := Department.GetDepartments()
	for ind, val := range depts {
		if jambite.CourseOfChoice == val.DepartmentName {
			if jambite.JambScore >= val.CutOffMark {
				Fname := jambite.FirstName
				Lname := jambite.LastName
				dept := jambite.CourseOfChoice
				matno := len(sch.StudentsList) + 207
				sch.StudentsList[ind-1] = Student.Student{Fname, Lname, dept, 0, matno, map[int]Course.Course{}, true}
			} else {
				fmt.Println("Jamb Score Too Low")
			}
		}
	}

}
func (p *Principal) UpdateStudentList(sch *School.School) {
	for ind, val := range sch.StudentsList {
		if val.PaidFeesInFull {
			p.StudentList[ind] = val
		} else {
			delete(sch.StudentsList, ind)
		}
	}
}

func (p *Principal) ExpelStudent(sch *School.School, name string) {
	for ind, val := range sch.StudentsList {
		if val.FirstName+" "+val.LastName == name {
			delete(sch.StudentsList, ind)
		}
	}
}
