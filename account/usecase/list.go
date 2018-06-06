package usecase

import (
	model "github.com/t-tashiro/yakiniku/account"
)

func (a *accountUsecase) List() ([]*model.Client, error) {
	accounts, err := a.accountRepository.List()
	if err != nil {
		return nil, err
	}
	res := formatCloentList(accounts)
	return res, nil
}
