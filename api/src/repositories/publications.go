package repositories

import (
	"database/sql"
	"gophernet/src/models"
)

// publications representa um repositorio de publicacoes
type publications struct {
	db *sql.DB
}

// NewPublicationsRepositories cria um repositorio de publicacoes
func NewPublicationsRepositories(db *sql.DB) *publications {
	return &publications{db}
}

// Create cria novas publicacoes de usuarios no banco de dados
func (repo publications) Create(publications models.Publications) (uint64, error) {
	statement, err := repo.db.Prepare(
		`
		INSERT INTO	publicacoes
			(titulo, conteudo, autorId)
		VALUES
			(?, ?, ?)
		`,
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(
		publications.Title,
		publications.Content,
		publications.AuthorID,
	)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

// FindByID retorna uma unica publicacao filtrada pelo ID
func (repo publications) FindByID(publicationsID uint64) (models.Publications, error) {
	row, err := repo.db.Query(
		`
		SELECT
			p.id, p.titulo, p.conteudo, p.autorId, u.nick, p.curtidas, p.criadaEm
		FROM
			publicacoes p
		INNER JOIN
			usuarios u
		ON
			p.autorId = u.id
		WHERE
			p.id = ?
		`,
		publicationsID,
	)
	if err != nil {
		return models.Publications{}, err
	}
	defer row.Close()

	var publication models.Publications

	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.AuthorNick,
			&publication.Likes,
			&publication.CreatedIn,
		); err != nil {
			return models.Publications{}, err
		}
	}

	return publication, nil
}

// Search
func (repo publications) Search() ([]models.Publications, error) {

	return nil, nil
}
