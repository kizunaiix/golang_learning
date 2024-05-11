package objs

type Student struct {
	Name string
	Age  int32
	Id   int64
}

type studentOption func(*Student)

func NewStudent(optlist ...studentOption) *Student {
	student := new(Student)
	for _, opt := range optlist {
		opt(student)
	}
	return student
}

func WithStudentName(n string) studentOption {
	return func(s *Student) { s.Name = n }

}

func WithStudentAge(age int32) studentOption {
	return func(s *Student) { s.Age = age }
}

func WithStudentId(id int64) studentOption {
	return func(s *Student) { s.Id = id }
}
