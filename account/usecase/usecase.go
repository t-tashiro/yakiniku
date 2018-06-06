package usecase

import (
	model "github.com/t-tashiro/yakiniku/account"
	"github.com/t-tashiro/yakiniku/account/repository"
)

type accountUsecase struct {
	accountRepository repository.AccountRepository
}

// NewAccountUsecase mount account usecase
func NewAccountUsecase(account repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepository: account,
	}
}

// AccountUsecase account usecase interface
type AccountUsecase interface {
	List() ([]*model.Client, error)
}

func formatCloentList(accounts []*model.Account) []*model.Client {
	res := []*model.Client{}
	for _, a := range accounts {
		account := formatClient(a)
		res = append(res, account)
	}
	return res
}

func formatClient(account *model.Account) *model.Client {
	res := &model.Client{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
	}
	return res
}
