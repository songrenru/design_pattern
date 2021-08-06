package mediator

import "fmt"

type ColleagueButton struct {
	name string
	Status bool
	m Mediator
}

func (b *ColleagueButton) SetMediator(m Mediator) {
	b.m = m
}

func (b *ColleagueButton) SetColleagueEnabled(status bool) {
	b.Status = status
	if status {
		fmt.Println("btn:%s 开启", b.name)
	} else {
		fmt.Println("btn:%s 关闭", b.name)
	}
}

type ColleagueTextfield struct {
	name string
	Content string
	m Mediator
}

func (b *ColleagueTextfield) SetMediator(m Mediator) {
	b.m = m
}

func (b *ColleagueTextfield) SetColleagueEnabled(status bool) {
	if status {
		fmt.Println("textfield:%s 开启", b.name)
	} else {
		fmt.Println("textfield:%s 关闭", b.name)
	}
}

type ColleagueCheckbox struct {
	name string
	Status bool
	m Mediator
}

func (b *ColleagueCheckbox) SetMediator(m Mediator) {
	b.m = m
}

func (b *ColleagueCheckbox) SetColleagueEnabled(status bool) {
	b.Status = status
	if status {
		fmt.Println("checkbox:%s 选中", b.name)
	} else {
		fmt.Println("checkbox:%s 未选中", b.name)
	}
}

type LoginFrame struct {
	CheckGuest, CheckLogin *ColleagueCheckbox
	TextUser, TextPswd *ColleagueTextfield
	BtnOk, BtnCancel *ColleagueButton
}

func NewLoginFrame() *LoginFrame {
	l := &LoginFrame{}
	l.CreateColleague()
	return l
}

func (l *LoginFrame) CreateColleague()  {
	l.CheckGuest = &ColleagueCheckbox{m: l, name: "guest"}
	l.CheckLogin = &ColleagueCheckbox{m: l, name: "login"}
	l.TextUser = &ColleagueTextfield{m: l, name: "user"}
	l.TextPswd = &ColleagueTextfield{m: l, name: "pswd"}
	l.BtnOk = &ColleagueButton{m: l, name: "ok"}
	l.BtnCancel = &ColleagueButton{m: l, name: "cancel"}
}

func (l *LoginFrame) ColleagueChanged()  {
	if l.CheckGuest.Status {
		l.TextUser.SetColleagueEnabled(false)
		l.TextPswd.SetColleagueEnabled(false)
		l.BtnOk.SetColleagueEnabled(true)
	} else {
		l.TextUser.SetColleagueEnabled(true)
		if len(l.TextUser.Content) > 0 {
			l.TextPswd.SetColleagueEnabled(true)
			if len(l.TextPswd.Content) > 0 {
				l.BtnOk.SetColleagueEnabled(true)
			} else {
				l.BtnOk.SetColleagueEnabled(false)
			}
		} else {
			l.TextPswd.SetColleagueEnabled(false)
			l.BtnOk.SetColleagueEnabled(false)
		}
	}
}