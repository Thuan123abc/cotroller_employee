package risefunc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ListExperience []Experience
var ListFresher []Fresher
var ListIntern []Intern

func (c Certificate) InputCerti() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("nhap ID bang cap\n")
	cerid, _ := consoleReader.ReadString('\n')
	cerid = strings.Replace(cerid, "\n", "", -1)
	ceridInt, _ := strconv.ParseInt(cerid, 10, 64)
	c.IDCertificate = ceridInt

	fmt.Println("nhap ten bang cap\n")
	cername, _ := consoleReader.ReadString('\n')
	cername = strings.Replace(cername, "\n", "", -1)
	c.NameCertificate = cername

	fmt.Println("nhap cap bac dat duoc cua bang cap\n")
	cerrank, _ := consoleReader.ReadString('\n')
	cerrank = strings.Replace(cerrank, "\n", "", -1)
	c.RankCertificate = cerrank

	fmt.Println("nhap ngay cap cua bang cap\n")
	cerdate, _ := consoleReader.ReadString('\n')
	cerdate = strings.Replace(cerdate, "\n", "", -1)
	c.DateCertificate = cerdate
}
func (e Employee) InputEmployee() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Nhap Id nhan vien\n")
	id, _ := consoleReader.ReadString('\n')
	id = strings.Replace(id, "\n", "", -1)
	idInt, _ := strconv.ParseInt(id, 10, 64)
	e.ID = idInt

	fmt.Println("Nhap ten nhan vien\n")
	name, _ := consoleReader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	e.Name = name

	fmt.Println("Nhap ngay sinh nhan vien\n")
	birth, _ := consoleReader.ReadString('\n')
	birth = strings.Replace(birth, "\n", "", -1)
	e.BirthDay = birth

	fmt.Println("Nhap sdt cua nhan vien\n")
	phone, _ := consoleReader.ReadString('\n')
	phone = strings.Replace(phone, "\n", "", -1)
	phoneInt, _ := strconv.ParseInt(phone, 10, 64)
	e.Phone = phoneInt

	fmt.Println("Nhap mail cua nhan vien\n")
	mail, _ := consoleReader.ReadString('\n')
	mail = strings.Replace(mail, "\n", "", -1)
	e.Email = mail

	e.Cert.InputCerti()
}

func (e Experience) InputExper() {
	e.InputEmployee()

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Nhap nam kinh nghiem cua experience\n")
	yers, _ := consoleReader.ReadString('\n')
	yers = strings.Replace(yers, "\n", "", -1)
	yersInt, _ := strconv.ParseInt(yers, 10, 64)
	e.ExpInYear = yersInt

	fmt.Println("Nhap ky nang cua experience\n")
	skill, _ := consoleReader.ReadString('\n')
	skill = strings.Replace(skill, "\n", "", -1)
	e.ProSkill = skill

	ListExperience = append(ListExperience, e)
}

func (f Fresher) InputFresher() {
	f.InputEmployee()

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("nhap ngay tot nghiep cua fresher\n")
	date, _ := consoleReader.ReadString('\n')
	date = strings.Replace(date, "\n", "", -1)
	f.GraduationDate = date

	fmt.Println("nhap xep loai tot nghiep cua fresher\n")
	rank, _ := consoleReader.ReadString('\n')
	rank = strings.Replace(rank, "\n", "", -1)
	f.GraduationRank = rank

	fmt.Println("nhap truong tot nghiep cua fresher\n")
	edu, _ := consoleReader.ReadString('\n')
	edu = strings.Replace(edu, "\n", "", -1)
	f.Education = edu

	ListFresher = append(ListFresher, f)
}

func (i Intern) InputIntern() {
	i.InputEmployee()

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("nhap chuyen nganh intern dang theo hoc\n")
	major, _ := consoleReader.ReadString('\n')
	major = strings.Replace(major, "\n", "", -1)
	i.Majors = major

	fmt.Println("nhap hoc ky intern dang theo hoc\n")
	semes, _ := consoleReader.ReadString('\n')
	semes = strings.Replace(semes, "\n", "", -1)
	semesInt, _ := strconv.ParseInt(semes, 10, 64)
	i.Semester = semesInt

	fmt.Println("nhap truong intern dang theo hoc\n")
	uni, _ := consoleReader.ReadString('\n')
	uni = strings.Replace(uni, "\n", "", -1)
	i.UniversityName = uni

	ListIntern = append(ListIntern, i)
}
