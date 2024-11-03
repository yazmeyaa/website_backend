package response

type ProjectsResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Img         string  `json:"img"`
	Href        *string `json:"href"`
	GithubUrl   *string `json:"githubUrl"`
	ImgUrl      *string `json:"imgUrl"`
}
