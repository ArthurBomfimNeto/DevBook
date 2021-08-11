package seguranca

import "golang.org/x/crypto/bcrypt"

//Cripytografar a formas de discrypitografar o hash não a como desfazer apenas compara-lo com uma string

//Hash recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost) // DefaultCost e um numero padrão para criação do hash
}

//VerificarSenha compara a senha e um hash e retorna se elas são iguais
func VerificarSenha(senhaString, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
