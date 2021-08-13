package utils

func Contains(arr []interface{}, value interface{}) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func ContainsString(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
