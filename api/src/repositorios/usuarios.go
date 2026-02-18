package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome,nick,email,senha) VALUES (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	insert, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	IDUsuario, err := insert.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(IDUsuario), nil
}
