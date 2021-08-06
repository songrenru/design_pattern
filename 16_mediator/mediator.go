package mediator

// 仲裁者
type Mediator interface {
	CreateColleague()
	ColleagueChanged()
}

// 同事
type Colleague interface {
	SetMediator(m Mediator)
	SetColleagueEnabled(status bool)
}
