package todo

type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}

	if s[i].Priority == s[j].Priority {
		return s[i].Position < s[j].Position
	}

	return s[i].Priority < s[j].Priority
}
