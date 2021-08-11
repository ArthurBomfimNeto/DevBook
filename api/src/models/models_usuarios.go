package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"` // omit empty caso for passar para json e o id estiver em branco elenão vai passar
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"Criadoem,omitempty"`
}

func (u *Usuario) Preparar(etapa string) error {
	if erro := u.validar(etapa); erro != nil {
		return erro
	}

	if erro := u.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (u *Usuario) validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("O nome é obrigatório!")
	}
	if u.Nick == "" {
		return errors.New("O nick é obrigatorio!")
	}
	if u.Email == "" {
		return errors.New("O email é obrigatorio!")
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("O email inserido é invalido, digite novamente!")
	}
	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("A senha é obrigatorio!")
	}
	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome) // TrimSpace remove apenas os espaços antes e depois da frase
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Senha = strings.TrimSpace(u.Senha)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(u.Senha)
		if erro != nil {
			return erro
		}
		u.Senha = string(senhaComHash)
	}
	return nil
}
