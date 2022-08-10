package risefunc

import "fmt"

type ShowMe interface {
	ShowMeExp(id int64) Experience
	ShowMeFresher(id int64) Fresher
	ShowMeIntern(id int64) Intern
}

func (e Experience) ShowMeExp(id int64) (Experience, error) {
	for i := 0; i < len(ListExperience); i++ {
		if id == ListExperience[i].ID {
			fmt.Println("nhan vien can tim la\n%v", ListExperience[i])
			break
		}
	}
	return e, nil
}

func (f Fresher) ShowMeFresher(id int64) (Fresher, error) {
	for i := 0; i < len(ListFresher); i++ {
		idFr := ListFresher[id].ID
		if id == idFr {
			fmt.Println("nhan vien can tim la\n%v", ListFresher[i])
			break
		}
	}
	return f, nil
}

func (i Intern) ShowMeIntern(id int64) (Intern, error) {
	for i := 0; i < len(ListIntern); i++ {
		idIn := ListIntern[id].ID
		if id == idIn {
			fmt.Println("nhan vien can tim la\n%v", ListIntern[i])
			break
		}
	}
	return i, nil
}
