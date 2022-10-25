package pointer

func Int(val int) *int {
	v := val
	return &v
}

func String(val string) *string {
	v := val
	return &v
}

func Bool(val bool) *bool {
	v := val
	return &v
}
