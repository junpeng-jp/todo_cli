package todo

func ByPriority(items []Item) func(i, j int) bool {

	return func(i, j int) bool {
		if items[i].Done != items[j].Done {
			return items[i].Done
		}

		if items[i].Priority == items[j].Priority {
			return items[i].CreatedAt < items[j].CreatedAt
		}

		return items[i].Priority < items[j].Priority
	}
}

func ByCreateTime(items []Item) func(i, j int) bool {

	return func(i, j int) bool {
		return items[i].CreatedAt < items[j].CreatedAt
	}
}
