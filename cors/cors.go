package cors

import "errors"

const (
	ORIGIN_ALLOWED     string = "allowed"
	ORIGIN_NOT_ALLOWED string = "not_allowed"
)

var (
	ErrRecordAlreadyExist error = errors.New("record with for this origin already exists")
	ErrNotFoundRecord     error = errors.New("cannot find record")
)

type CORSRecord struct {
	OriginAllowed  bool
	AllowedMethods []string
	AllowedHeaders []string
}
