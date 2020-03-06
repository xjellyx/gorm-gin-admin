package model

import (
	"github.com/go-redis/redis"
	base "github.com/olongfen/userDetail"
	"sync"
	"time"
)

func initCorn() {
	t := time.NewTicker(time.Hour)
	for {
		select {
		case <-t.C:
			func() {
				var (
					err    error
					uidArr []string
					rdb    *redis.Client
				)
				if err = Database.Model(&UserOnline{}).Select(`uid`).Where("is_online = ?", true).Scan(&uidArr).
					Error; err != nil {
					LogModel.Warnln("[cron user offline] err: ", err)
					return
				}
				if len(uidArr) == 0 {
					return
				}
				if rdb, err = GetRedisClient(); err != nil {
					LogModel.Warnln("[cron user offline] GetRedisClient err: ", err)
					return
				}
				defer rdb.Close()
				wg := sync.WaitGroup{}
				for _, v := range uidArr {
					if _d, _err := rdb.Exists("cache_token" + v).Result(); _err != nil {
						continue
					} else if _d == 0 {
						d := false
						wg.Add(1)
						go func() {
							defer wg.Done()
							_, _ = PubUserOnlineUpdate(v, &base.FormUserOnline{IsOnline: &d})
						}()
					}
				}
				wg.Wait()
			}()
		}
	}
}
