package account

import "challenge/internal/ports"

type Handler struct {
	AccountService ports.AccountService
}
