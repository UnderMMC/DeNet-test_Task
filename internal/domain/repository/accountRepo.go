package repository

import "DeNet-test_Task/internal/domain/entity"

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) GetAccount(address string) (*entity.Account, error) {
	return nil, nil
}

func (r *AccountRepository) GetERC20Balance(address, tokenAddress string) (*entity.ERC20Balance, error) {
	return nil, nil
}
