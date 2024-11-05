package request

type AddCorsOriginRequest struct {
	Origin         string   `json:"origin" validate:"min=6,max=500,required"`
	AllowedMethods []string `json:"allowedMethods"`
	AllowedHeaders []string `json:"allowedHeaders"`
}
