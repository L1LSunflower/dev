package facade

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type accRepo map[string]*Account

var (
	accountRepository = accRepo{}
	userRole          = "user"

	errAccountDoesNotExist = errors.New("account does not exist by this id")
)

type Facade struct {
	account    *Account
	repository accRepo
	validator  *Validator
	auth       *Authenticator
	logger     *Logger
}

type Repository interface {
	StoreAccount(id, username, password string, age int)
	GetByID(id string) (*Account, error)
}

func (a *accRepo) StoreAccount(id, username, password string, age int) {
	account := &Account{}
	account.SetID(id).SetUsername(username).SetPassword(password).SetAge(age).SetRole(userRole)
	accountRepository[id] = account.GetAccount()
}

func (a *accRepo) GetByID(id string) (*Account, error) {
	if account, ok := accountRepository[id]; ok {
		return account, nil
	}
	return nil, errAccountDoesNotExist
}

type Account struct {
	id       string
	username string
	password string
	age      int
	role     string
}

type AccountBuilder interface {
	SetID(id string) AccountBuilder
	SetUsername(username string) AccountBuilder
	SetPassword(password string) AccountBuilder
	SetAge(age int) AccountBuilder
	SetRole(role string) AccountBuilder
}

type AccountInterface interface {
	AccountBuilder
	ID() string
	Username() string
	Password() string
	Age() int
	Role() string
	GetAccount() *Account
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) Username() string {
	return a.username
}

func (a *Account) Password() string {
	return a.password
}

func (a *Account) Age() int {
	return a.age
}

func (a *Account) Role() string {
	return a.role
}

func (a *Account) SetID(id string) AccountBuilder {
	a.id = id
	return a
}

func (a *Account) SetUsername(username string) AccountBuilder {
	a.username = username
	return a
}

func (a *Account) SetPassword(password string) AccountBuilder {
	a.password = password
	return a
}

func (a *Account) SetAge(age int) AccountBuilder {
	a.age = age
	return a
}

func (a *Account) SetRole(role string) AccountBuilder {
	a.role = role
	return a
}

func (a *Account) GetAccount() *Account {
	return a
}

type Validator struct{}

func (v *Validator) Validate(values map[string]any) error {
	for key, value := range values {
		switch val := value.(type) {
		case string:
			if len(val) <= 0 {
				return fmt.Errorf("wrong value for %s", key)
			}
		case int:
			if val <= 0 {
				return fmt.Errorf("wrong value for %s", key)
			}
		default:
			return fmt.Errorf("unexpected type")
		}
	}
	return nil
}

type Authenticator struct{}

func (a *Authenticator) Auth(id, password string, repository Repository) error {
	account, err := repository.GetByID(id)
	if err != nil {
		return err
	}

	if password != account.Password() {
		return fmt.Errorf("password dont match")
	}

	return nil
}

type Logger struct{}

func (l *Logger) Log(typeLog string, message any) {
	log.Printf("[TIME:%d]\tLOG_LEVEL:%s\tMESSAGE:%v\n", time.Now().Unix(), typeLog, message)
}
