package models

type CadastroFisio struct {
	Nome  string `json:"nome_fisioterapeuta"`
	Email string `json:"email_fisioterapeuta"`
	Senha string `json:"senha_fisioterapeuta"`
}