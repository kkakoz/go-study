package reflects

import "testing"

func TestProvider(t *testing.T) {
	Provider(NewStudent, NewTeacher)
}

type Student struct {
	name string
}

func NewStudent(name string) *Student {
	return &Student{name: name}
}

type Teacher struct {
	student Student
}

func NewTeacher(student Student) *Teacher {
	return &Teacher{student: student}
}