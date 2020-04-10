package models

import (
	"errors"
	"fmt"
	"time"
	"github.com/cicdi-go/sso/src/utils"
	"github.com/xormplus/xorm"
	"github.com/go-redis/redis"
)

type ActiveRecod interface {
	GetDb() (e *xorm.Engine, err error)
	TableName() string
}

type Base struct {
}

func (u *Base) TableName() string {
	return utils.Config.TablePrefix + "user"
}

func (u *Base) GetDb() (e *xorm.Engine, err error) {
	var found bool
	if e, found = utils.Engin.GetXormEngin("postgres"); !found {
		err = errors.New("Database default is not found!")
	}
	return
}

func SetRedis(captcha string,email string) (e *redis.Client, found bool) {

	if e, found = utils.RedisClient.Get("redis"); !found {
		err := errors.New("Database default is not found!")
		fmt.Println(err)
	}

	err := e.Set(email, captcha, 300*time.Second).Err()
    if err != nil {
            fmt.Println(err)
           //panic(err)
       }
	return
}

func GetRedis() (e *redis.Client,err error) {
    var found bool
	if e, found = utils.RedisClient.Get("redis"); !found {
		err := errors.New("redis default is not found!")
		fmt.Println(err)
	}
	return
}


