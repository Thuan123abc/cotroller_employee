package risefunc

type Certificate struct {
	IDCertificate   int64
	NameCertificate string
	RankCertificate string
	DateCertificate string
}
type Employee struct {
	ID       int64
	Name     string
	BirthDay string
	Phone    int64
	Email    string
	Cert     Certificate
}
type Experience struct {
	Employee
	ExpInYear int64
	ProSkill  string
}
type Fresher struct {
	Employee
	GraduationDate string
	GraduationRank string
	Education      string
}
type Intern struct {
	Employee
	Majors         string
	Semester       int64
	UniversityName string
}
