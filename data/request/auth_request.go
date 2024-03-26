package request

type AuthCredentails struct {
	Username string `validate:"required,min=1,max=64" json:"username"`
	Password string `validate:"required,min=1,max=255" json:"password"`
}
