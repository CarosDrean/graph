package ordered

type Int struct{}

func NewOrderedInt() Int {
	return Int{}
}

func (o Int) IsLeft(s, t any) bool {
	return s.(int) > t.(int)
}

func (o Int) IsEqual(s, t any) bool {
	return s.(int) == t.(int)
}
