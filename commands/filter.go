package commands

func Filter(tasks []Task, predicate func(t Task) bool) []Task {
	var result []Task
	for _, t := range tasks {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}
