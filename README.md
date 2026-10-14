# 🐹 GopherNet

Rede social inspirada no Twitter, com **API REST em Go** e um **front-end web também em Go**. Os usuários se cadastram, fazem login com JWT, seguem outros usuários, publicam posts e curtem publicações.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)

---

## ✨ Funcionalidades

- **Usuários:** cadastro, busca, edição, exclusão e troca de senha
- **Autenticação:** login com **JWT** (expira em 6h) e senhas protegidas com **bcrypt**
- **Seguidores:** seguir, deixar de seguir, listar seguidores e quem o usuário segue
- **Publicações:** criar, listar, editar, excluir, curtir e descurtir
- **Middlewares:** log de requisições e verificação do token nas rotas protegidas

## 🏗️ Arquitetura

O projeto tem duas aplicações:

```
gophernet/
├── api/                 # API REST (porta 5000)
│   ├── main.go
│   ├── authentication/  # geração e validação de JWT
│   └── src/
│       ├── config/       # variáveis de ambiente (.env)
│       ├── controllers/  # handlers HTTP
│       ├── middlewares/  # logger e autenticação
│       ├── models/       # entidades e validações
│       ├── repositories/ # acesso ao banco (SQL puro)
│       ├── router/       # definição das rotas (gorilla/mux)
│       ├── responses/    # respostas JSON e erros padronizados
│       └── security/     # hash de senha com bcrypt
├── webapp/              # front-end em Go + templates HTML (porta 3000)
└── sql/                 # script de criação do banco
```

A API segue uma separação em camadas: **router → middleware → controller → repository → MySQL**.

## 🔌 Endpoints

| Método | Rota | Auth |
|---|---|:---:|
| POST | `/login` | |
| POST | `/usuarios` | |
| GET | `/usuarios?usuario={nome ou nick}` | 🔒 |
| GET · PUT · DELETE | `/usuarios/{id}` | 🔒 |
| POST | `/usuarios/{id}/seguir` · `/parar-de-seguir` | 🔒 |
| GET | `/usuarios/{id}/seguidores` · `/seguindo` | 🔒 |
| POST | `/usuarios/{id}/atualizar-senha` | 🔒 |
| GET | `/usuarios/{id}/publicacoes` | 🔒 |
| POST · GET | `/publicacoes` | 🔒 |
| GET · PUT · DELETE | `/publicacoes/{id}` | 🔒 |
| POST | `/publicacoes/{id}/curtir` · `/tirar-curtida` | 🔒 |

🔒 = enviar o header `Authorization: Bearer <token>`

## 🚀 Como rodar

**Pré-requisitos:** Go 1.25+ e MySQL.

**1. Banco de dados**

```bash
mysql -u root -p < sql/sql.sql
```

**2. API**

Crie o arquivo `api/.env`:

```env
DB_USER=root
DB_PASSWORD=sua_senha
DB_NAME=gophernet
API_PORT=5000
SECRET_KEY=uma_chave_secreta_grande
```

```bash
cd api
go run main.go
```

**3. WebApp**

```bash
cd webapp
go run main.go
```

Acesse **http://localhost:3000**.

## 🛣️ Próximos passos

- [ ] Telas de feed, perfil e publicações no webapp
- [ ] Docker Compose para subir API + MySQL com um comando
- [ ] Testes automatizados
- [ ] Documentação da API com Swagger

---

Feito por [Cícero Guilherme](https://github.com/CiceroGGS)
