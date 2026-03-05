CREATE DATABASE gophernet;

USE gophernet;

---------------------------------------------------------------------------------------------------------------

CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE ,
    email  VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);

---------------------------------------------------------------------------------------------------------------

CREATE TABLE seguidores (
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,

    seguidor_id INT NOT NULL,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
);

---------------------------------------------------------------------------------------------------------------

---------------------------------------------------------------------------------------------------------------


CREATE TABLE publicacoes (
        id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,

    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id) REFERENCES usuarios(id) ON DELETE CASCADE,

    curtidas INT DEFAULT 0,
    criadaEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);

---------------------------------------------------------------------------------------------------------------

SELECT
    s.usuario_id,
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
	u.id = s.seguidor_id;


---------------------------------------------------------------------------------------------------------------

SELECT
    s.usuario_id,
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
	u.id = s.seguidor_id;

---------------------------------------------------------------------------------------------------------------

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
    s.seguidor_id = 1;

---------------------------------------------------------------------------------------------------------------
