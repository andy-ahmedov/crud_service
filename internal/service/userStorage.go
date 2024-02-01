package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/andy-ahmedov/crud_service/internal/domain"
	"github.com/golang-jwt/jwt"
)

// добавляем новое поле sessionRepo SessionRepository
type Users struct {
	repo   UserStorage
	hasher PasswordHasher

	hmacSecret []byte
	tokenTtl   time.Duration
}

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type UserStorage interface {
	CreateUser(ctx context.Context, inp domain.User) error
	GetByCredential(ctx context.Context, email string, passwords string) (domain.User, error)
}

// добавляем новый интерфейс SessionRepository 
// Create(ctx context.Context, token domain.RefreshSession) error
// Get(ctx ..... , token string) (domain.RefreshSession, error)


// также добавляем новое поле в NewUsers
func NewUsers(repo UserStorage, hasher PasswordHasher, secret []byte, ttl time.Duration) *Users {
	return &Users{
		repo:       repo,
		hasher:     hasher,
		hmacSecret: secret,
		tokenTtl:   ttl,
	}

}

func (u *Users) SignUp(ctx context.Context, inp domain.SignUpInput) error {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:         inp.Name,
		Email:        inp.Email,
		Password:     password,
		RegisteredAt: time.Now(),
	}

	return u.repo.CreateUser(ctx, user)
}


// изменить signIn функцию, чтобы она возвращала accesT refreshT и err
func (u *Users) SignIn(ctx context.Context, inp domain.SignInInput) (string, error) {
	password, err := u.hasher.Hash(inp.Password)
	if err != nil {
		return "", err
	}

	getIt, err := u.repo.GetByCredential(ctx, inp.Email, password)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return "", domain.ErrUserNotFound
		}
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(u.tokenTtl).Unix(),
		Subject:   strconv.Itoa(int(getIt.ID)),
	})

	// возвращаем метод generateTokens(ctx, user.ID)
	return token.SignedString(u.hmacSecret)
}

func (u *Users) ParseToken(ctx context.Context, token string) (int64, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return u.hmacSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subject := claims["sub"].(string)
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, errors.New("invalid subject")
	}
	
	return int64(id), nil
}


// generateTokens(ctx , userID int64) (string string error)
// перенести логику jwt.NewWithClaims и 
// ..., ... := t.SignedString()
// после чего используем функцию newRefreshTokens для генерации рефрештокена
// используем метод create сущности sessionRepo для создания новой рефрешСессии, вставив в поля userID, refreshToken и время окончания через 30 дней. Проверяем на ошибки
// возвращаем аксесс рефреш и ошибку


//  newRefreshToken() (string, err)
// создаем слайс байтов с емкостью 32
// создаем newSource с помощью одноименной функции в библиотеке ранд, положив туда нынешнее время в формате юникс
// полсе чего используем newSource для создания нового структуры ранд используя функцию New в библе ранд
// используем метод Read из переменной выше в которой кладем созданный слайс, проверяем на ошибки
// возвращаем строку в формате %x

// RefreshTokens(ctx, refreshToken string) (string, string, error)
// используем метод Get сущности sessionRepo для получения сессии. Проверяем на ошибки
// проверяем если в полученной сессии поле ExpiresAt преобразованная в unix меньше нынешнего времени того же формата, возвращаем ошибку из domain ErrRefreshTokenExpired
// возвращем метод generateTokens()