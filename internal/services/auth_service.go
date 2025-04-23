package services

import (
	"bookstore_go_v1/internal/models"
	"bookstore_go_v1/internal/repositories"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
    Register(username, password string) error
    Login(username, password string) (string, string, error)
    RefreshToken(refreshToken string) (string, string, error)
    Logout(userID string) error
}

type authService struct {
    userRepo repositories.UserRepository
    secret   string
    authExp  int
    refreshExp int
}

func NewAuthService(
    userRepo repositories.UserRepository,
    secret string,
    authExp int,
    refreshExp int,
) AuthService {
    return &authService{
        userRepo:   userRepo,
        secret:     secret,
        authExp:    authExp,
        refreshExp: refreshExp,
    }
}

func (s *authService) Register(username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &models.User{
        Username: username,
        Password: string(hashedPassword),
    }

    return s.userRepo.CreateUser(user)
}

func (s *authService) Login(username, password string) (string, string, error) {
    user, err := s.userRepo.FindByUsername(username)
    if err != nil {
        return "", "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", "", errors.New("invalid credentials")
    }

    authClaims := jwt.MapClaims{
        "sub": strconv.Itoa(user.ID),
        "exp": time.Now().Add(time.Second * time.Duration(s.authExp)).Unix(),
    }

    authJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
    authToken, _ := authJwt.SignedString([]byte(s.secret))

    refreshClaims := jwt.MapClaims{
        "sub": strconv.Itoa(user.ID),
        "exp": time.Now().Add(time.Second * time.Duration(s.refreshExp)).Unix(),
    }

    refreshJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    refreshToken, _ := refreshJwt.SignedString([]byte(s.secret))

    expTime := time.Now().Add(time.Second * time.Duration(s.refreshExp)).Unix()
    s.userRepo.UpdateRefreshToken(strconv.Itoa(user.ID), refreshToken, expTime)

    return authToken, refreshToken, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, string, error) {
    user, err := s.userRepo.FindByRefreshToken(refreshToken)
    if err != nil {
        return "", "", errors.New("invalid refresh token")
    }

    if user.RefreshTokenExp.Unix() < time.Now().Unix() {
        return "", "", errors.New("refresh token expired")
    }

    authClaims := jwt.MapClaims{
        "sub": strconv.Itoa(user.ID),
        "exp": time.Now().Add(time.Second * time.Duration(s.authExp)).Unix(),
    }

    authJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
    authToken, _ := authJwt.SignedString([]byte(s.secret))

    newRefreshToken, newRefreshExp := generateRefreshToken(s.secret, s.refreshExp, strconv.Itoa(user.ID))
    s.userRepo.UpdateRefreshToken(strconv.Itoa(user.ID), newRefreshToken, newRefreshExp)

    return authToken, newRefreshToken, nil
}

func (s *authService) Logout(userID string) error {
    return s.userRepo.UpdateRefreshToken(userID, "", 0)
}

func generateRefreshToken(secret string, exp int, userID string) (string, int64) {
    expTime := time.Now().Add(time.Second * time.Duration(exp)).Unix()
    claims := jwt.MapClaims{
        "sub": userID,
        "exp": expTime,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, _ := token.SignedString([]byte(secret))
    return signedToken, expTime
}