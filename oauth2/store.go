package oauth2

type ClientStore interface {
	SetInfo(key string, val interface{}) error
	GetInfo(key string) (interface{}, error)
}
