package proxy

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrCreateAccount   = errors.New("account already exists")
)

type AccountProxy interface {
	GetAccount(id string) (*Account, error)
}

type Proxy struct {
	dbRepo    DBRepository
	cacheRepo CacheRepository
}

func (p *Proxy) GetAccount(id string) (*Account, error) {
	acc, err := p.cacheRepo.GetAccount(id)
	if err != nil && !errors.Is(err, ErrAccountNotFound) {
		if acc, err = p.dbRepo.GetAccount(id); err != nil && !errors.Is(err, ErrAccountNotFound) {
			return nil, ErrAccountNotFound
		}

		if err = p.cacheRepo.StoreAccount(acc); err != nil && !errors.Is(err, ErrCreateAccount) {
			return nil, ErrCreateAccount
		}

		return acc, nil
	}

	return nil, ErrAccountNotFound
}

type DBRepository struct {
	tempDB map[string]*Account
}

func (r *DBRepository) StoreAccount(account *Account) error {
	if _, err := r.GetAccount(account.id); err != nil && !errors.Is(err, ErrAccountNotFound) {
		return ErrCreateAccount
	}
	r.tempDB[account.id] = account
	return nil
}

func (r *DBRepository) GetAccount(id string) (*Account, error) {
	if account, ok := r.tempDB[id]; ok {
		return account, nil
	}
	return nil, ErrAccountNotFound
}

type CacheRepository struct {
	tempCacheDB map[string]*Account
}

func (r *CacheRepository) StoreAccount(account *Account) error {
	if _, err := r.GetAccount(account.id); err != nil && !errors.Is(err, ErrAccountNotFound) {
		return ErrCreateAccount
	}
	r.tempCacheDB[account.id] = account
	return nil
}

func (r *CacheRepository) GetAccount(id string) (*Account, error) {
	if account, ok := r.tempCacheDB[id]; ok {
		return account, nil
	}
	return nil, ErrAccountNotFound
}

type Account struct {
	id       string
	username string
	password string
}
