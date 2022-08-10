package database

type Certificate struct {
	IDCertificate   int64 `grom:"primaryKey"`
	NameCertificate string
	RankCertificate string
	DateCertificate string
}
type Employee struct {
	ID             int64 `grom:"primaryKey"`
	Name           string
	BirthDay       string
	Phone          int64
	Email          string
	Cert           Certificate `grom:"ForeignKey:IDCertificate"`
	ExpInYear      int64
	ProSkill       string
	GraduationDate string
	GraduationRank string
	Education      string
	Majors         string
	Semester       int64
	UniversityName string
}

func (c Certificate) TableName() string {
	return "certificate"
}
func (e Employee) TableName() string {
	return "employee"
}
