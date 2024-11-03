package schemas

const (
	SCHEMA_TYPE_STRING = "string"
	SCHEMA_TYPE_NUMBER = "number"
)

type FieldSchema struct {
	Type     string `json:"type"`
	Nullable bool   `json:"nullable"`
	IsArray  bool   `json:"isArray"`
	Readonly bool   `json:"isReadonly"`
}

type ProjectSchema struct {
	ID          FieldSchema `json:"id"`
	Name        FieldSchema `json:"name"`
	Description FieldSchema `json:"description"`
	Href        FieldSchema `json:"href"`
	Img         FieldSchema `json:"img"`
	GithubUrl   FieldSchema `json:"githubUrl"`
	ImgUrl      FieldSchema `json:"imgUrl"`
}

func NewProjectSchema() *ProjectSchema {
	return &ProjectSchema{
		ID:          FieldSchema{Type: SCHEMA_TYPE_NUMBER, Nullable: false, IsArray: false, Readonly: true},
		Name:        FieldSchema{Type: SCHEMA_TYPE_STRING, Nullable: false, IsArray: false, Readonly: false},
		Description: FieldSchema{Type: SCHEMA_TYPE_STRING, Nullable: false, IsArray: false, Readonly: false},
		Href:        FieldSchema{Type: SCHEMA_TYPE_STRING, Nullable: true, IsArray: false, Readonly: false},
		Img:         FieldSchema{Type: SCHEMA_TYPE_STRING, Nullable: false, IsArray: false, Readonly: false},
		GithubUrl:   FieldSchema{Type: SCHEMA_TYPE_STRING, Nullable: true, IsArray: false, Readonly: false},
		ImgUrl:      FieldSchema{Type: SCHEMA_TYPE_STRING, Nullable: false, IsArray: false, Readonly: false},
	}
}

type StaticFileSchema struct {
	ID       FieldSchema `json:"id"`
	Name     FieldSchema `json:"name"`
	FileName FieldSchema `json:"fileName"`
}

func NewStaticFileSchema() *StaticFileSchema {
	return &StaticFileSchema{
		ID:       FieldSchema{Type: SCHEMA_TYPE_NUMBER, Nullable: false, IsArray: false, Readonly: true},
		Name:     FieldSchema{Type: SCHEMA_TYPE_STRING, Nullable: false, IsArray: false, Readonly: false},
		FileName: FieldSchema{Type: SCHEMA_TYPE_STRING, Nullable: true, IsArray: false, Readonly: false},
	}
}
