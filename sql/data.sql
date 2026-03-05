INSERT INTO usuarios
    (nome, nick, email, senha)
VALUES
    ("Usuario 1", "usuario_1", "emailusuario111@com.br", "$2a$10$UWJOCepn.brzwnosOhmkqeGrLlglWkfZ0XLdXgA5Vrv.aTvb0i73O"),
    ("Usuario 2", "usuario_2", "emailusuario222@com.br", "$2a$10$UWJOCepn.brzwnosOhmkqeGrLlglWkfZ0XLdXgA5Vrv.aTvb0i73O"),
    ("Usuario 3", "usuario_3", "emailusuario333@com.br", "$2a$10$UWJOCepn.brzwnosOhmkqeGrLlglWkfZ0XLdXgA5Vrv.aTvb0i73O");

---------------------------------------------------------------------------------------------------------------

INSERT INTO seguidores
    (usuario_id, seguidor_id)
VALUES
    (1, 2),
    (3, 1),
    (1, 4);

---------------------------------------------------------------------------------------------------------------

