package account

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func (h Handler) HandleRequest(ctx context.Context) error {
	account, err := h.AccountService.Get(ctx)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return err
	}

	fmt.Println(account)

	h.AccountService.Create(account)

	h.AccountService.Send(account)

	return nil
}
