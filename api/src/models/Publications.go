package models

import "time"

type Publications struct {
	ID         uint64    `json:"publicacaoId,omitempty"`
	Title      string    `json:"titulo,omitempty"`
	Content    string    `json:"conteudo,omitempty"`
	AuthorID   uint64    `json:"autorId,omitempty"`
	AuthorNick string    `json:"autorNick,omitempty"`
	Likes      uint64    `json:"curtidas"`
	CreatedIn  time.Time `json:"criadaEm"`
}
