package models

type User struct {
	IDFisioterapeuta string `id_fisioterapeuta`
	Nome             string `json:"nome_fisioterapeuta"`
	Email            string `json:"email_fisioterapeuta"`
	Senha            string `json:"senha_fisioterapeuta"`
}
