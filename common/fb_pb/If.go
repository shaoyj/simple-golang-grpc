package fb_pb

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
}
