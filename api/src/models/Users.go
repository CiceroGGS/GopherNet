package models

import (
	"errors"
	"gophernet/src/security"
	"net/mail"
	"strings"
	"time"
)

// User representa um usuario utilizando a aplicacao
type Users struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"senha,omitempty"`
	CreatedIn time.Time `json:"criadoEm,omitempty"`
}

func (user *Users) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *Users) validate(step string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco.")
	}
	if user.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco.")
	}
	if user.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco.")
	}
	if step == "register" && user.Password == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco.")
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return errors.New("E-mail em formato invalido")
	}

	return nil
}

func (user *Users) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
