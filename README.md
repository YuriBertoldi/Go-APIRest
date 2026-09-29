# 🌐 Go-APIRest — API REST com Gorilla Mux

API REST em **Go** com **Gorilla Mux**, **GORM** e **PostgreSQL**, com CRUD completo de personalidades. Projeto de estudo de 2023.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)

## 📬 Endpoints

| Método | Rota | Ação |
|---|---|---|
| `GET` | `/api/personalidades` | Lista todas |
| `GET` | `/api/personalidades/{id}` | Busca por ID |
| `POST` | `/api/personalidades` | Cria |
| `PUT` | `/api/personalidades/{id}` | Edita |
| `DELETE` | `/api/personalidades/{id}` | Remove |

Inclui middleware de `Content-Type` JSON e CORS liberado.

## 🚀 Como rodar

```bash
cd Banco && docker compose up -d   # PostgreSQL + script inicial
cd .. && go run main.go
```

API disponível em **http://localhost:8000**.

---

Desenvolvido por **Yuri Bertoldi** — [LinkedIn](https://www.linkedin.com/in/yuri-bulh%C3%B5es-bertoldi-b62459180/)
