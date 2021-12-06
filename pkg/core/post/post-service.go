package post

import "fmt"

type PostService struct {
}

func (r *PostService) Get(id int) (*Post, error) {
	post := Post{
		Id:   id,
		Text: "Post " + fmt.Sprint(id),
	}
	return &post, nil
}

func (r *PostService) Create(data *CreatePostData) (*Post, error) {
	fmt.Println("creating post: " + data.Text)
	post := Post{
		Id:   1337,
		Text: data.Text,
	}
	return &post, nil
}
