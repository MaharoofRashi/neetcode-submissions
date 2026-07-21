func numUniqueEmails(emails []string) int {
    seen := make(map[string]bool)

	for _, email := range emails {
		parts := strings.SplitN(email, "@", 2)
		local, domain := parts[0], parts[1]
		if idx := strings.Index(local, "+"); idx != -1 {
			local = local[:idx]
		}
		local = strings.ReplaceAll(local, ".", "")
		seen[local+"@"+domain] = true
	}

	return len(seen)
}