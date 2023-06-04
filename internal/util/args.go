package util

func GroupBySize(slice []string, groupSize int) [][]string {
	var groups [][]string
	var currentGroup []string

	for _, element := range slice {
		currentGroup = append(currentGroup, element)
		if len(currentGroup) == groupSize {
			groups = append(groups, currentGroup)
			currentGroup = []string{}
		}
	}

	if len(currentGroup) > 0 {
		groups = append(groups, currentGroup)
	}

	return groups
}
