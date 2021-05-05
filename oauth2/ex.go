package oauth2

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/url"
)

type ClientDB struct {
	store map[string]interface{}
}

func (c ClientDB) SetInfo(key string, val interface{}) error {
	c.store[key] = val
	return nil
}

func (c ClientDB) GetInfo(key string) (interface{}, error) {
	v := c.store[key]
	return v, nil
}

type TokenRedis struct {
	store map[string]interface{}
}

func (t TokenRedis) SetInfo(key string, val interface{}) error {
	t.store[key] = val
	return nil
}

func (t TokenRedis) GetInfo(key string) (interface{}, error) {
	v := t.store[key]
	return v, nil
}

func (t TokenRedis) Refresh(key string, expire uint64) error {
	t.store[key] = key
	return nil
}

func (t TokenRedis) AccessTokenGenerate() string {
	panic("implement me")
}

func ex() {
	router := gin.Default()
	router.Any("/authorize", func(c *gin.Context) {
		session := sessions.Default(c)
		inface := session.Get("ReturnUri")
		var form url.Values
		if val, ok := inface.(url.Values); ok {
			form = val
		}
		session.Delete("ReturnUri")
		session.Save()
	})
}
