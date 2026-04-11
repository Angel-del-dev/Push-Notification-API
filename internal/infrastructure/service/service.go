package service

func NewService[T any]() *T {
	return new(T)
}
