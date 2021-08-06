package chain_of_responsibility

import "fmt"

// Support 这里我只定义了一个resolve()方法，是否support的判定，都交给具体了 // 还是需要HaveRight()方法
type Support interface {
	HaveRight(t *Trouble) bool
	Resolve(t *Trouble) bool
}

// CommonSupport
type CommonSupport struct {
	Support // 嵌套一个interface，为了给他包装一层setNext()方法
	next *CommonSupport // CommonSupport即实现了support，也提供了一个通用的setNext()方法
}

func (s *CommonSupport) SetNext(next *CommonSupport) *CommonSupport {
	s.next = next
	return next // 链式操作
}

func (s *CommonSupport) HaveRight(t *Trouble) bool {
	return s.Support.HaveRight(t)
}

func (s *CommonSupport) Resolve(t *Trouble) bool { // todo
	if s.HaveRight(t) {
		return s.Support.Resolve(t)
	} else if s.next != nil {
		return s.next.Resolve(t)
	} else {
		//fmt.Println("无人解决，完犊子了！")
		return false
	}
}

// LimitSupport 只实现support，再把LimitSupport嵌入CommonSupport使用
type LimitSupport struct {
	limit int
}

func NewLimitSupport(limit int) *CommonSupport {
	return &CommonSupport{
		Support: &LimitSupport{limit: limit},
	}
}

func (s *LimitSupport) HaveRight(t *Trouble) bool {
	return t.Num < s.limit
}

func (s *LimitSupport) Resolve(t *Trouble) bool {
	fmt.Printf("limit 解决了问题, t num = %d \n", t.Num)
	return true
}

// OddSupport
type OddSupport struct {}

func NewOddSupport() *CommonSupport {
	return &CommonSupport{
		Support: &OddSupport{},
	}
}

func (s *OddSupport) HaveRight(t *Trouble) bool {
	return t.Num & 1 == 1
}

func (s *OddSupport) Resolve(t *Trouble) bool {
	fmt.Printf("odd 解决了问题, t num = %d \n", t.Num)
	return true
}

// SepecialSupport
type SepecialSupport struct {
	num int
}

func NewSepecialSupport(num int) *CommonSupport {
	return &CommonSupport{
		Support: &SepecialSupport{num: num},
	}
}

func (s *SepecialSupport) HaveRight(t *Trouble) bool {
	return t.Num == s.num
}

func (s *SepecialSupport) Resolve(t *Trouble) bool {
	fmt.Printf("sepecial 解决了问题, t num = %d \n", t.Num)
	return true
}

// Trouble 不是必须的，只是为了象征下待解决的问题
type Trouble struct {
	Num int
}

func NewTrouble(num int) *Trouble {
	return &Trouble{Num: num}
}


