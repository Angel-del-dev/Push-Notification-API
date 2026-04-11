package domaincreator

func Create[T any]() *T {
	return new(T)
}
