package utils

import (
	"errors"
	"server/global"
	"server/model/request"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	AccessTokenSecret  []byte
	RefreshTokenSecret []byte
}

var (
	TokenExpired     = errors.New("token expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		AccessTokenSecret:  []byte(global.Config.JWT.AccessTokenSecret),
		RefreshTokenSecret: []byte(global.Config.JWT.RefreshTokenSecret),
	}
}

// CreateAccessClaims 创建 Access Token 的 Claims，包含基本信息和过期时间等
func (j *JWT) CreateAccessClaims(baseClaims request.BaseClaims) request.JwtCustomClaims {
	ep, _ := ParseDuration(global.Config.JWT.AccessTokenExpiryTime)
	claims := request.JwtCustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"TAP"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    global.Config.JWT.Issuer,
		},
	}
	return claims
}

// CreateAccessToken 创建 Access Token，通过 Claims 生成 JWT Token
func (j *JWT) CreateAccessToken(claims request.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //创建新的JWT token
	return token.SignedString(j.AccessTokenSecret)             //返回签名后的token
}

// CreateRefreshClaims 创建 Refresh Token 的 Claims，包含用户信息和过期时间等
func (j *JWT) CreateRefreshClaims(baseClaims request.BaseClaims) request.JwtCustomRefreshClaims {
	ep, _ := ParseDuration(global.Config.JWT.RefreshTokenExpiryTime)
	return request.JwtCustomRefreshClaims{
		UserID: baseClaims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"TAP"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			Issuer:    global.Config.JWT.Issuer,
		},
	}
}

// CreateRefreshToken 创建 Refresh Token，通过 Claims 生成 JWT Token
func (j *JWT) CreateRefreshToken(claims request.JwtCustomRefreshClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //创建新的JWT token
	return token.SignedString(j.RefreshTokenSecret)            //返回签名后的token
}

// ParseAccessToken 解析 Access Token，验证 Token 并返回 Claims 信息
func (j *JWT) ParseAccessToken(token string) (*request.JwtCustomClaims, error) {
	// 实现解析 Access Token 的逻辑
	claims, err := j.parseToken(token, &request.JwtCustomClaims{}, j.AccessTokenSecret)
	if err != nil {
		return nil, err
	}
	if customClaims, ok := claims.(*request.JwtCustomClaims); ok {
		return customClaims, nil
	}
	return nil, TokenInvalid // Token 无效，返回错误
}

func (j *JWT) ParseRefreshToken(token string) (*request.JwtCustomRefreshClaims, error) {
	claims, err := j.parseToken(token, &request.JwtCustomRefreshClaims{}, j.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}
	// 	claims (interface{})
	//   ↓
	// 	尝试转换成 *request.JwtCustomRefreshClaims
	//   ↓
	// 	转换成功 → ok=true, refreshClaims 是具体类型
	// 	转换失败 → ok=false, refreshClaims 是该类型的零值
	if refreshClaims, ok := claims.(*request.JwtCustomRefreshClaims); ok {
		return refreshClaims, nil
	}
	return nil, TokenInvalid // Token 无效，返回错误
}

// parseToken 通用的 Token 解析方法，验证 Token 是否有效并返回 Claims 信息
func (j *JWT) parseToken(tokenString string, claims jwt.Claims, secretKey interface{}) (interface{}, error) {
	// 用于解析并验证 JWT（JSON Web Token）
	// tokenString：JWT token 字符串
	// claims：用于接收解析后数据的结构体（必须实现 Claims 接口）
	// keyFunc：验证签名的密钥函数
	token, err := jwt.ParseWithClaims(tokenString, claims, func(toke *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		if ve, ok := err.(jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, TokenMalformed // Token 格式错误
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, TokenExpired // Token 已过期
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, TokenNotValidYet // Token 还未生效
			default:
				return nil, TokenInvalid // 其他错误返回 Token 无效
			}

		}
		return nil, TokenInvalid // 默认返回 Token 无效错误
	}
	if token.Valid { // 如果 Token 验证通过，返回 Claims
		return token.Claims, nil
	}
	return nil, TokenInvalid // Token 无效，返回错误
}
