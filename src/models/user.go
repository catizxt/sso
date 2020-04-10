package models

import (
	"github.com/cicdi-go/sso/src/utils"
	"time"
	"errors"
	//"log"
	"fmt"
	"github.com/cicdi-go/jwt"
)

type User struct {
	*Base              `xorm:"-"`
	Id                 int64  `json:"id"`
	Username           string `xorm:"varchar(100) notnull index default ''" json:"username"`
	Role               string `xorm:"varchar(100) default ''" json:"role"`
	Email              string `xorm:"varchar(50) default ''" json:"email"`
	Status             int    `xorm:"SMALLINT default 1" json:"status"`
	AuthKey            string `xorm:"varchar(32) default ''" json:"-"`
	PasswordHash       string `xorm:"varchar(255) default ''" json:"-"`
	PasswordResetToken string `xorm:"varchar(255) default ''" json:"-"`
	password           string `xorm:"-"`
	CreatedAt          int    `xorm:"created" json:"created_at"`
	UpdatedAt          int    `xorm:"updated" json:"updated_at"`
}

func (u *User) TableName() string {
	return utils.Config.TablePrefix + "user"
}

/*func init() {
	u := new(User)
	if e, err := u.GetDb(); err != nil {
		log.Println(err)
	} else {
		err := e.Sync2(u)
		if err != nil {
			log.Println(err)
		}
	}
}*/

func (u *User) Update(captcha string) (err error) {
	engine, err := u.GetDb()
	if err != nil {
		return
	}

    user := new(User)

    //判断邮箱号是否注册过账号
    has, err := engine.Table("public.user").Where("email = ?", u.Email).Get(user)
    if err != nil {
        return err
    }
    if has != true {
        err := errors.New("不存在该用户")
        return err
    }

    //判断注册码是否正确
    re , err := GetRedis()
    if err!= nil{
        return
    }
    val, err := re.Get(u.Email).Result()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("redis获取："+val)
    fmt.Println("用户传入："+captcha)
    if val != captcha {
        err := errors.New("验证码输入错误")
        return err
    }

    //fmt.Println("更新这里不对，没有获得id")
	affected, err := engine.Id(user.Id).Update(u) //更新用户
	if err != nil {
	    //.Println("update出错")
	    fmt.Println(affected)
		return err
	}
	return
}


func (u *User) Insert(captcha string) (err error) {
	engine, err := u.GetDb()
	if err != nil {
		return
	}

    //判断邮箱号是否注册过账号
    has, err := engine.Table("public.user").Where("email = ?", u.Email).Exist()

    if err != nil {
        fmt.Println(has)
        fmt.Println(err)
        return err
    }

    if has == true {
        err := errors.New("该用户已注册过账号")
        return err
    }

    //判断注册码是否正确
    re , err := GetRedis()
    if err!= nil{
        return
    }

    val, err := re.Get(u.Email).Result()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("redis获取："+val)
    fmt.Println("用户传入："+captcha)

    if val != captcha {
        err := errors.New("验证码输入错误")
        return err
    }

	id, err := engine.Insert(u)
	if err != nil {
		return err
	}
	u.Id = id
	return
}

func (u *User) SetPassword(value string) {
	u.password = value
	u.generateAuthKey()
	u.PasswordHash, _ = utils.SetPassword(u.password, u.AuthKey)
}

func (u *User) GetPasswordHash(p string) string {
	passwordHash, err := utils.SetPassword(p, u.AuthKey)
	if err != nil {
		return ""
	}
	return passwordHash
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) generateAuthKey() {
	u.AuthKey = utils.GenerateRandomKey()
}

func (u *User) Verify(p string) bool {
	engine, err := u.GetDb()
	if err != nil {
		return false
	}
	engine.Where("email = ?", u.Email).Get(u)
	//log.Println(u.GetPasswordHash(p))
	//log.Println(u.PasswordHash)
	return u.GetPasswordHash(p) == u.PasswordHash
}

// 生成jwt
func (u *User) GenerateToken() (token string, expire time.Time, err error) {
	algorithm :=  jwt.HmacSha256("cicdi")
	claims := jwt.NewClaim()
	expire = time.Now().Add(time.Second*time.Duration(utils.Config.Expire*1000000))
	claims.Set("Username", u.Email)
	claims.Set("Role", u.Role)
	claims.SetTime("exp", expire)
	token, err = algorithm.Encode(claims)
	return token, expire, err
}
