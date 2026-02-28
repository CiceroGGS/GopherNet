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
		return errors.New("usuário não")
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
