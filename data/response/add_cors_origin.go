package response

type AddCorsOriginResponse struct {
	Origin         string   `json:"origin"`
	OriginAllowed  bool     `json:"allowed"`
	AllowedMethods []string `json:"allowedMethods"`
	AllowedHeaders []string `json:"allowedHeaders"`
}
