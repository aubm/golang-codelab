package basics

type MemberOfTheFamily struct {
	Name     string
	Children []MemberOfTheFamily
}

func NumberOfDescendants(member MemberOfTheFamily) int {
}
