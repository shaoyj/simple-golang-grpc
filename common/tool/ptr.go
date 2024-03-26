package tool

func ToPtr[T any](data T) *T {
	return &data
}
