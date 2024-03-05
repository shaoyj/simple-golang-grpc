package fb_pb

func ToPtr[T any](data T) *T {
	return &data
}
