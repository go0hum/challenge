package domain

type EmailInfo struct {
	from    string
	to      []string
	subject string
	body    string
}
