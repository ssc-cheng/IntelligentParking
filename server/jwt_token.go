package server

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
	"intelligent_parking/utils/errors"
)

// 用户发送token时的结构体数据
type jwtCustomClaims struct {
	jwt.StandardClaims
	// 追加自己需要的信息
	//Uid uint `json:"uid"`
	//Admin bool `json:"admin"`
}

func NewJwtToken() *jwtCustomClaims {
	return new(jwtCustomClaims)
}


// 生成token
func (j *jwtCustomClaims)CreateToken(SecretKey []byte, user_id ,user_account string) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 2).Unix()),  // jwt的过期时间
			Id:    user_id,
			Issuer:user_account,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}





// ParseToken 是jwt校验
func (j *jwtCustomClaims)ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var (
		token *jwt.Token
	)
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}

//获取头部携带token并解析
func (j *jwtCustomClaims)Gtoken(headertoken string) (err error) {

	// 获取请求头的信息，做初步判断
	if headertoken == "" {
		errors.New("没有请求头信息，请携带请求头数据")
		return
	}
	// 解析token
	_, err = j.ParseToken(headertoken, []byte(beego.AppConfig.String("secretKey")))
	if err != nil {
		errors.Wrap(err,"该token已被篡改")
		return
	}

	// 检验jwt是否过期或被篡改
	if err = j.Valid();err != nil{
		return
	}

	return
}
