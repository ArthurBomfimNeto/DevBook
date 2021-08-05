package modelos

import (
	"errors"
	"strings"
	"time"
)

type Cordenada struct {
	Longitude string `json: "longitude"`
	Latitude  string `json: "latitude"`
}

type Usuario struct {
	ID        uint64    `json: "id,omitempty"` // omit empty caso for passar para json e o id estiver em branco elenão vai passar
	Nome      string    `json: "nome,omitempty"`
	Nick      string    `json: "nick,omitempty"`
	Email     string    `json: "email,omitempty"`
	Senha     string    `json: "senha,omitempty"`
	CriadoEm  time.Time `json: "Criadoem,omitempty"`
	Cordenada Cordenada `json: "cordenada"`
}

func (u *Usuario) Preparar() error {
	if erro := u.validar(); erro != nil {
		return erro
	}

	u.formatar()
	return nil
}

func (u *Usuario) validar() error {
	if u.Nome == "" {
		return errors.New("O nome é obrigatório!")
	}
	if u.Nick == "" {
		return errors.New("O nick é obrigatorio!")
	}
	if u.Email == "" {
		return errors.New("O email é obrigatorio!")
	}
	if u.Senha == "" {
		return errors.New("O senha é obrigatorio!")
	}
	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome) // TrimSpace remove apenas os espaços antes e depois da frase
	u.Email = strings.TrimSpace(u.Email)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Senha = strings.TrimSpace(u.Senha)
}
