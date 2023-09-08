package captcha

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type RedisStore struct {
	Ctx       context.Context
	KeyPrefix string
}

func (s *RedisStore) Set(key string, value string) error {

	expireTime := g.Cfg("captcha").MustGet(s.Ctx, "captcha.expire_time").Int64()

	error := g.Redis("cache").SetEX(s.Ctx, s.KeyPrefix+key, value, expireTime*60)

	return error
}

func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key

	val, err := g.Redis("cache").Get(s.Ctx, key)

	if err != nil {
		return ""
	}

	if clear {
		g.Redis("cache").Del(s.Ctx, key)
	}
	return val.String()
}

func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
