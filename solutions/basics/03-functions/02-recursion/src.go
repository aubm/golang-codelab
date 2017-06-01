package basics

type MemberOfTheFamily struct {
	Name     string
	Children []MemberOfTheFamily
}

func NumberOfDescendants(member MemberOfTheFamily) int {
	nbDescendants := len(member.Children)
	for _, child := range member.Children {
		nbDescendants += NumberOfDescendants(child)
	}
	return nbDescendants
}
