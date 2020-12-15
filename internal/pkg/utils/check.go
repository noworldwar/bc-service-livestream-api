package utils

func CheckJsonDataContent(data interface{}, vals ...string) string {
	for _, v := range vals {
		m := data.(map[string]interface{})
		if m[v] == nil {
			return v
		}
	}
	return ""
}
