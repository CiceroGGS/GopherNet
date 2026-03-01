package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"gophernet/src/models"
)

type users struct {
	db *sql.DB
}

// NewUsersRepositories cria e retorna uma nova instância do repositório de usuários com a conexão ao banco de dados injetada.
func NewUsersRepositories(db *sql.DB) *users {
	return &users{db}
}

// CreateUsersDB insere um novo usuário no banco de dados e retorna o ID gerado.
func (repo users) Create(user models.Users) (uint64, error) {
	statement, err := repo.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(
		&user.Name,
		&user.Nick,
		&user.Email,
		&user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	user.ID = uint64(userID)

	return user.ID, nil
}

// GetUsersDB busca e retorna todos os usuários cujo nome ou nick correspondam ao filtro informado.
func (repo users) Search(nameOrNick string) ([]models.Users, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repo.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nick LIKE ? OR nome LIKE ?",
		nameOrNick,
		nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.Users

	for rows.Next() {
		var user models.Users
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByIDFromDB busca e retorna um único usuário no banco de dados pelo seu ID.
func (repo users) FindByID(ID uint64) (models.Users, error) {
	row, err := repo.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", ID)
	if err != nil {
		return models.Users{}, err
	}
	defer row.Close()

	var user models.Users

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return models.Users{}, err
		}
	}

	return user, nil
}

// UpdateUserInDB atualiza o nome, nick e email de um usuário no banco de dados pelo seu ID.
func (repo users) Update(user models.Users, ID uint64) error {
	statement, err := repo.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		&user.Name,
		&user.Nick,
		&user.Email,
		ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserInDB remove permanentemente um usuário do banco de dados pelo seu ID, retornando erro caso não seja encontrado.
func (repo users) Delete(ID uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	exec, err := statement.Exec(ID)
	if err != nil {
		return err
	}

	rows, _ := exec.RowsAffected()
	if rows == 0 {
		return errors.New("usuário não localizado")
	}

	return nil
}

// SearchByEmail busca usuario pelo seu E-mail e retorna seu id e senha com hash.
func (repo users) SearchByEmail(email string) (models.Users, error) {
	row, err := repo.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if err != nil {
		return models.Users{}, err
	}
	defer row.Close()

	var user models.Users

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.Users{}, err
		}
	}

	return user, nil
}

// Follow registra um novo relacionamento de seguidor entre usuários.
func (repo users) Follow(userID, followerID uint64) error {
	statement, err := repo.db.Prepare("INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow remove o relacionamento de seguidor entre usuários.
func (repo users) Unfollow(userID, followerID uint64) error {
	statement, err := repo.db.Prepare(
		`
		DELETE FROM
			seguidores
		WHERE
			usuario_id = ?
		AND
			seguidor_id = ?
		`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// FindFollowers bysca os seguidores existente de um usuario na aplicacao
func (repo users) FindFollowers(userID uint64) ([]models.Users, error) {
	rows, err := repo.db.Query(
		`
		SELECT
			u.id,
			u.nome,
			u.nick,
			u.email,
			u.criadoEm
		FROM
			usuarios u
		INNER JOIN
			seguidores s
		ON
			u.id = s.seguidor_id
		WHERE
			s.usuario_id = ?
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.Users

	for rows.Next() {
		var user models.Users
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// FindFollowing busca os usuarios que estao sendo seguidos por algum outro usuario
func (repo users) FindFollowing(userID uint64) ([]models.Users, error) {
	rows, err := repo.db.Query(
		`
		SELECT
			u.id,
			u.nome,
			u.nick,
			u.email,
			u.criadoEm
		FROM
			usuarios u
		INNER JOIN
			seguidores s
		ON
			u.id = s.usuario_id
		WHERE
			s.seguidor_id = ?
	`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	var users []models.Users

	for rows.Next() {
		var user models.Users
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// UpdatePasswordByID atualiza a senha de um usuario na aplicacao filtrando pelo ID
func (repo users) UpdatePassword(userID uint64, newPassword string) error {
	statement, err := repo.db.Prepare(
		`
		UPDATE
			usuarios
		SET
			senha = ?
		WHERE
			id = ?
		`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(&newPassword, userID); err != nil {
		return err
	}

	return nil
}

func (repo users) FindPassword(userID uint64) (string, error) {
	row, err := repo.db.Query(
		`
		SELECT
			senha
		FROM
			usuarios
		WHERE
			id = ?
		`,
		userID,
	)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var password models.Password
	currentPassword := password.CurrentPassword

	if row.Next() {
		if err = row.Scan(
			&currentPassword,
		); err != nil {
			return "", err
		}
	}

	return currentPassword, nil
}
