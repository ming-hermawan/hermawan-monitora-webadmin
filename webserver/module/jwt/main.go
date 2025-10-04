package jwt

import (
    "log"
    "errors"
    "fmt"
    "time"
    "net/http"
    "github.com/golang-jwt/jwt"
    "hermawan-monitora/module/hmonenv"
    "hermawan-monitora/webserver/module/httpresponse"
)


// PUBLIC

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte(hmonenv.GetJwtSignatureKey())

func GetJwtToken(w http.ResponseWriter, username string) string {
    claims := jwt.MapClaims{
      "username": username,
      "exp": time.Now().Add(time.Hour * 1).Unix(),
    }
    token := jwt.NewWithClaims(
      JWT_SIGNING_METHOD,
      claims)
    out, err := token.SignedString(JWT_SIGNATURE_KEY)
    if err != nil {
        httpresponse.ErrResponseForInvalidToken(
          w,
          fmt.Sprintf("Token %s Error <%s>",
                      username,
                      err.Error()))
        return ""
    }
    return out
}

func GetUsernameFromToken(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Signing Method Invalid")
        } else if method != JWT_SIGNING_METHOD {
            return nil, fmt.Errorf("Signing Method Invalid")
        }
        return JWT_SIGNATURE_KEY, nil
    })
    if err != nil {
        return "", err
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return "", err
    }
    username := claims["username"].(string)
    exp := int64(claims["exp"].(float64))
    log.Println(fmt.Printf("NOW = %d\n", time.Now().Unix()))
    log.Println(fmt.Printf("EXPIRED = %d\n", exp))
    if time.Now().Unix() < exp {
        return username, nil
    } else {
        return "", errors.New("Token Expired")
    }
}
