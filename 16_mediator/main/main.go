package main

import mediator "github.com/songrenru/design_attern/16_mediator"

func main() {
	l := mediator.NewLoginFrame()

	//l.CheckGuest.SetColleagueEnabled(true) // 选中guest
	l.CheckLogin.SetColleagueEnabled(true) // 选中login
	l.TextUser.Content = "eason"


	l.ColleagueChanged()
}
