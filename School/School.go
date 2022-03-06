package School

import (
	"schoolmodel/Course"
	"schoolmodel/Student"
	"schoolmodel/Teacher"
)

type School struct {
	SchoolName   string
	TeachersList []Teacher.Teacher
	StudentsList []Student.Student
	CourseList   []Course.Course
}
