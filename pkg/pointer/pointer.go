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

func IntVal(val *int) int {
	if val == nil {
		return 0
	}
	return *val
}

func StringVal(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

func BoolVal(val *bool) bool {
	if val == nil {
		return false
	}
	return *val
}
