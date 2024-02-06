package ordered

type String struct{}

func NewOrderedString() String {
	return String{}
}

func (o String) IsLeft(s, t any) bool {
	return s.(string) > t.(string)
}

func (o String) IsEqual(s, t any) bool {
	return s.(string) == t.(string)
}
