package ordered

type Ordered interface {
	IsLeft(s, t any) bool
	IsEqual(s, t any) bool
}
