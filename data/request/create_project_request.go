package request

type CreateProjectRequest struct {
	Name        string  `validate:"required,min=1,max=200" json:"name"`
	Description string  `validate:"required,min=1,max=500" json:"description"`
	Img         string  `validate:"required,min=1,max=500" json:"img"`
	Href        *string `json:"href"`
	GithubUrl   *string `json:"githubUrl"`
	ImgUrl      string  `validate:"min=1,max=500" json:"imgUrl"`
}
