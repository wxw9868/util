package common

//// 从redis取JWT
//func GetRedisJWT(key string) (result string, err error) {
//	result, err = redis.GetRedisClient().Get(key).Result()
//	return
//}
//
//// JWT存入redis并设置过期时间
//func SetRedisJWT(key string, value string) error {
//	timer := time.Duration(config.GetConfig().JWT.ExpiresTime) * time.Second
//	err := redis.GetRedisClient().Set(key, value, timer).Err()
//	return err
//}
