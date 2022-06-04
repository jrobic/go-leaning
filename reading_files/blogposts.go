package readingfiles

import (
	"io/fs"
)

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(filesystem, f)
		if err != nil {
			//todo: needs clarification, should we totally fail if one file fails? or just ignore?
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
