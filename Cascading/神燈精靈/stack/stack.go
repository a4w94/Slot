package stack

type Stack struct {
	list []interface{}
}

func New() *Stack {
	s := new(Stack)
	s.list = make([]interface{}, 0)
	return s
}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

//取出slcie 最尾端元素
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return tmp

}

func (s *Stack) Push(element interface{}) {
	s.list = append(s.list, element)
}

func (s *Stack) ToSlice() []interface{} {
	return s.list
}
