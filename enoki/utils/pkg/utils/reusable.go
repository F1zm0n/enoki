package utilities

func MapContains(m1, m2 map[string]bool) bool {
	for k, v := range m2 {
		if val, ok := m1[k]; !ok || val != v {
			return false
		}
	}
	return true
}
