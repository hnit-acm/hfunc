package services

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/hnit-acm/hfunc/hapi/hsessions"
	"log"
)

var (
	_SessionStore sessions.Store
)

func SessionStore() sessions.Store {
	if _SessionStore == nil {
		s, err := hsessions.NewRedisStoreClusterCli(RedisCluster())
		if err != nil {
			log.Panicln(err)
		}
		_SessionStore = s
	}
	return _SessionStore
}
