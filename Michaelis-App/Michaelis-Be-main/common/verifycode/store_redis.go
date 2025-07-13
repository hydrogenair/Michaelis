package verifycode

import (
	rds "LingXi/common/redis"
	"time"
)

type RedisStore struct {
	RedisClient *rds.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string, expire int) error {
	ExpireTime := time.Minute * time.Duration(expire)
	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// 检查验证码
func (s *RedisStore) Verify(id, answer string) bool {
	v := s.Get(id, false)
	if v == answer {
		s.Get(id, true)
	}
	return v == answer
}
