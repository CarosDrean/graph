package ordered

type Ordered interface {
	IsValidType(data any) bool
	IsLeft(s, t any) bool
	IsEqual(s, t any) bool
}
