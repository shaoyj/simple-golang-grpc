package pb

func ToPtr[T any](data T) *T {
	return &data
}