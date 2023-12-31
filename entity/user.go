package entity

import (
	"final-project/infrastructure/config"
	"final-project/pkg/errs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role" binding:"required,oneof= member"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")

func (u *User) HashPassword() errs.MessageErr {
	salt := 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)

	if err != nil {
		return errs.NewInternalServerError("SOMETHING WENT WRONG")
	}

	u.Password = string(bs)

	return nil
}
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":        u.Id,
		"full_name": u.FullName,
		"role":      u.Role,
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := config.AppConfig().SecretKey

	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func (u *User) parseToken(tokenString string) (*jwt.Token, errs.MessageErr) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}

		secretKey := config.AppConfig().SecretKey

		return []byte(secretKey), nil
	})

	if err != nil {

		return nil, invalidTokenErr
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	if id, ok := claim["id"].(float64); !ok {
		return invalidTokenErr
	} else {
		u.Id = int(id)
	}

	if fullName, ok := claim["full_name"].(string); !ok {
		return invalidTokenErr
	} else {
		u.FullName = fullName
	}

	if role, ok := claim["role"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Role = role
	}

	return nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return errs.NewUnauthenticatedError("token should be Bearer")
	}

	splitToken := strings.Fields(bearerToken)

	if len(splitToken) != 2 {
		return errs.NewUnauthenticatedError("invalid token")
	}

	tokenString := splitToken[1]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errs.NewUnauthenticatedError("invalid token" + err.Error())
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}
