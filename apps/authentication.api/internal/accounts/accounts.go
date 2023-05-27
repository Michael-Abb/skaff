package accounts

type accountService struct{}

type AccountsService interface{}

func NewAccountService() *accountService {
	return &accountService{}
}
