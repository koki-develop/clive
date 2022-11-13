package cmd

func ptrString(v string) *string {
	return &v
}

func contains(slice []string, r string) bool {
	for _, l := range slice {
		if l == r {
			return true
		}
	}
	return false
}
