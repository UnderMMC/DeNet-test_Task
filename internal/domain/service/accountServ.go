package service

import (
	"DeNet-test_Task/internal/domain/entity"
)

type AccountRepository interface {
	GetAccount(address string) (*entity.Account, error)
	GetERC20Balance(address, tokenAddress string) (*entity.ERC20Balance, error)
}

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(repo AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (r *AccountService) GetAccount(address, signature string) (*entity.Account, error) {
	// account, err := r.GetAccount()
	return nil, nil
}

func (r *AccountService) GetAccounts(addresses []string, tokenAddress string) ([]entity.ERC20Balance, error) {
	return nil, nil
}
