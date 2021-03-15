package oauth2

type TokenStore interface {
	SetInfo(key string, val interface{}) error
	GetInfo(key string) (interface{}, error)
	Refresh(key string, expire uint64) error
	AccessTokenGenerate() string
}
