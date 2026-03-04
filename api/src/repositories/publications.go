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

// Search busca todas as publicacoes de todos usuarios que o usuario da aplicao segue
func (repo publications) Search(userID uint64) ([]models.Publications, error) {
	rows, err := repo.db.Query(
		`
		SELECT DISTINCT
			p.id, p.titulo, p.conteudo, p.autorId,
			u.nick, p.curtidas, p.criadaEm
		FROM
			publicacoes p
		INNER JOIN
			usuarios u ON p.autorId = u.id
		LEFT JOIN
			seguidores s ON p.autorId = s.usuario_id
		WHERE
			p.autorId = ? OR s.seguidor_id = ?
		`,
		userID,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publications

	for rows.Next() {
		var publication models.Publications

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.AuthorNick,
			&publication.Likes,
			&publication.CreatedIn,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Update atualiza os dados de uma publicação específica no banco.
func (repo publications) Update(publicationID uint64, publication models.Publications) error {
	statement, err := repo.db.Prepare(
		`
		UPDATE
			publicacoes
		SET
			titulo = ?,
			conteudo = ?
		WHERE
			id = ?
		`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}

	return nil
}

// Delete remove permanentemente uma publicação do banco de dados com base no ID informado.
func (repo publications) Delete(publicationsID uint64) error {
	statement, err := repo.db.Prepare(
		`
		DELETE FROM
			publicacoes
		WHERE
			id = ?
		`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicationsID); err != nil {
		return err
	}

	return nil
}

// SearchByUser busca todas as publicacoes de um usuario da aplicacao
func (repo publications) SearchByUser(userID uint64) ([]models.Publications, error) {
	rows, err := repo.db.Query(
		`
		SELECT
			p.id, p.titulo, p.conteudo, p.autorId,
			u.nick, p.curtidas, p.criadaEm
		FROM
			publicacoes p
		INNER JOIN
			usuarios u
		ON
			 u.id = p.autorId
		WHERE
			p.autorId = ?
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publications

	for rows.Next() {
		var publication models.Publications

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.AuthorNick,
			&publication.Likes,
			&publication.CreatedIn,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}
