package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/schema"
	user "github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/user"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/utils"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var mySigningKey = []byte(os.Getenv("AUTH_SECRET_KEY"))

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


type AuthService struct {
	pb.UnimplementedAuthServiceServer
	AuthClient pb.AuthServiceClient
	ReadDB         *ent.Client
	WriteDB        *ent.Client
}

func (c *AuthService) Register(ctx context.Context, input *pb.UserRegisterInput) (*pb.RegisterPayload, error) {
	// check if username existed
	existedUser, err := c.ReadDB.User.Query().Where(user.Username(input.Username)).Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	if existedUser > 0 {
		return nil, fmt.Errorf("username existed")
	}

	// create hashed password
	// Tham khảo tại https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
	pwd := []byte(input.Password)
	hashedPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("failed hashing password: %w", err)
	}

	// Create new user
	now := time.Now()
	wut := time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, now.Location())
	slt := time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, now.Location())

	userCreated, err := c.WriteDB.User.Create().
		SetUsername(input.Username).
		SetPassword(string(hashedPwd)).
		SetWakeUpTime(wut).
		SetSleepTime(slt).
		SetGender(schema.Gender(0)).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating new user: %w", err)
	}

	// Create the Claims
	// Tham khảo https://pkg.go.dev/github.com/golang-jwt/jwt#example-NewWithClaims-CustomClaimsType
	claims := MyCustomClaims{
		fmt.Sprint(userCreated.Username),
		jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return nil, fmt.Errorf("failed signing token: %w", err)
	}

	return &pb.RegisterPayload{
		User:  utils.MapUserRecordToUserProto(userCreated),
		Token: &pb.TokenPayloadDto{
			ExpiresIn:   utils.IntToOptionalInt32(15000),
			AccessToken: &ss,
		},
	}, nil
}