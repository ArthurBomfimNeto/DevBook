package modelos

import "time"

type Usuario struct {
	ID       uint64    `json: "id,omitempty"` // omit empty caso for passar para json e o id estiver em branco elenão vai passar
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "senha,omitempty"`
	CriadoEm time.Time `json: "Criadoem,omitempty"`
}
