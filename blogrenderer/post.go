package blogrenderer

import "strings"

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
