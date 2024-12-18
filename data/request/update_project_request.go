package request

type UpdateProjectRequest struct {
	ID          int    `validate:"required" json:"id"`
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Description string `validate:"required,min=1,max=500" json:"description"`
	Img         string `validate:"required,min=1,max=500" json:"img"`
	Href        string `validate:"min=1,max=500" json:"href"`
	GithubUrl   string `validate:"min=1,max=500" json:"githubUrl"`
	ImgUrl      string `validate:"min=1,max=500" json:"imgUrl"`
}
