package middleware

import (
	"gin-gorm/utils"
	"gin-gorm/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func SetToken(username string) (string, int) {
    ExpiresAtTime := time.Now().Add(time.Hour * 10)
    claims := MyClaims{
        username,
        jwt.StandardClaims{
            ExpiresAt: ExpiresAtTime.Unix(),
            Issuer:    "gin-gorm",
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(JwtKey)
    if err !=nil {
        return "",errmsg.ERROR
    }
    return tokenString, errmsg.SUCCESS
}


func ParseToken(tokenString string) (*MyClaims, int) {
    token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })
    if err != nil {
        return nil, errmsg.ERROR_TOKEN_EXIST
    }
    if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
        return claims, errmsg.SUCCESS
    }
    return nil, errmsg.ERROR_TOKEN_EXIST
}

func JwtToken() gin.HandlerFunc{
    return func(c *gin.Context) {
        code := errmsg.SUCCESS
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            code = errmsg.ERROR_TOKEN_EXIST
            c.JSON(http.StatusOK,gin.H{
                "code": code,
                "msg":  errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        checkToken := strings.SplitN(tokenString, " ", 2)
        if len(checkToken) != 2 || checkToken[0] != "Bearer"{
            code = errmsg.ERROR_TOKEN_WRONG
            c.JSON(http.StatusOK,gin.H{
                "code": code,
                "msg":  errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        claims, code := ParseToken(checkToken[1])
        if code != errmsg.SUCCESS {
            code = errmsg.ERROR_TOKEN_WRONG
            c.JSON(http.StatusOK,gin.H{
                "code": code,
                "msg":  errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        if claims.ExpiresAt < time.Now().Unix() {
            code = errmsg.ERROR_TOKEN_RUNTIME
            c.JSON(http.StatusOK,gin.H{
                "code": code,
                "msg":  errmsg.GetErrMsg(code),
            })
            c.Abort()
            return
        }
        c.Set("username", claims.UserName)
        c.Next()
    }
}