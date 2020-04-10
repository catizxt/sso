package handler

import (
	"context"
	"github.com/micro/go-log"
	"fmt"
	sso "github.com/cicdi-go/sso/proto/sso"
	models "github.com/cicdi-go/sso/src/models"
	"errors"
	ssoutils "github.com/cicdi-go/sso/src/utils"
	"github.com/astaxie/beego/utils"
)

type Sso struct{}

// Call is a single request handler called via client.Auth or the generated client code
func (e *Sso) Token(ctx context.Context, req *sso.AuthRequest, rsp *sso.AuthResponse) error {
	rsp.Status = 0
	u := models.User{
		Email: req.Username,
	}
	if !u.Verify(req.Password) {
		err := errors.New("账号或者密码错误")
		return err
	}
	token, expire, err := u.GenerateToken()
	if err != nil {
		return err
	}
	rsp.Expire = expire.Unix()
	rsp.Token = token

	rsp.Status = 1
	return nil
}

// Call is a single request handler called via client.Register or the generated client code
func (e *Sso) Register(ctx context.Context, req *sso.RegisterRequest, rsp *sso.RegisterResponse) error {

	if req.Mobile == "" {
	    return errors.New("邮箱号不能为空")
	}

	if req.Captcha == "" {
    		return errors.New("验证码不能为空")
    }

	user := models.User{
	    Email : req.Mobile,
		Role : "user",
		Status:   1,
	}
	rsp.Status = 0
	user.SetPassword(req.Password)
	if err := user.Insert(req.Captcha); err != nil {
		return err
	}
	rsp.Status = int64(user.Status)
	return nil
}

//email,password,captcha
func (e *Sso) ForgetPassword(ctx context.Context, req *sso.PasswordRequest, rsp *sso.PasswordResponse) error {

	if req.Mobile == "" {
	    return errors.New("邮箱号不能为空")
	}

	if req.Captcha == "" {
    		return errors.New("验证码不能为空")
    }

	user := models.User{
	    Email : req.Mobile,
		Role : "user",
		Status:   1,
	}
	rsp.Status = 0
	user.SetPassword(req.Password)

	if err := user.Update(req.Captcha); err != nil {
		return err
	}
	rsp.Status = int64(user.Status)
	return nil
}


//发送邮件验证函数
func (e * Sso) SendEmail(ctx context.Context, req *sso.EmailRequest, rsp *sso.EmailResponse) (error){
    	// 发送激活邮件
    	config := `{"username":"778439511@qq.com","password":"uwjrktowawmjbajb","host":"smtp.qq.com","port":587}`
    	temail := utils.NewEMail(config)

    	//指定收件人邮箱地址
    	temail.To = []string{req.Email}

    	//指定发件人的邮箱地址
    	temail.From = "778439511@qq.com"

    	//指定邮件的标题
    	temail.Subject = "网站注册用户验证码"

    	//指定邮件内容
        captcha := ssoutils.GenValidateCode(6)
    	temail.HTML = "复制该验证码到大数据自学网站账号注册处，验证码："+captcha

        models.SetRedis(captcha,req.Email)

        fmt.Println(temail.To)

    	err := temail.Send()
    	if err != nil {
    		rsp.Result = false
    		return nil
    	}
    	rsp.Result = true
    	return nil
}


/*func (e * Sso) Captcha(ctx context.Context, req *sso.CaptchaRequest, rsp *sso.CaptchaResponse) error {
	id, data, err := utils.CaptchaGenerate(req.Type, req.Length)
	if err != nil {
		return err
	}
	rsp.Id = id
	rsp.Data = data
	return nil
}*/

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Sso) Stream(ctx context.Context, req *sso.StreamingRequest, stream sso.Sso_StreamStream) error {
	log.Logf("Received Sso.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&sso.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Sso) PingPong(ctx context.Context, stream sso.Sso_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&sso.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

