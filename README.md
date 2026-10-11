# Mandarin App

> Uma aplicação fullstack interativa para o ensino e prática da língua mandarim.

---

## 📌 Sobre o Projeto

O **Mandarin App** é uma plataforma desenvolvida para auxiliar estudantes no aprendizado de mandarim, oferecendo recursos práticos para o estudo de ideogramas (Hanzi), pinyin, vocabulário e pronúncia.

---

## 🛠️ Tecnologias Utilizadas

### **Frontend**
* [React.js](https://reactjs.org/)
* TypeScript
* Tailwind CSS

### **Backend**
* Golang

### **Banco de Dados**
* PostgreSQL

---

## 📁 Estrutura do Projeto

```text
mandarin-app/
├── .github/
│   └── workflows/
│       └── ci.yml
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── database/
│   │   └── db.go
│   ├── entity/
│   │   ├── ideogram.go
│   │   └── repository.go
│   ├── handlers/
│   │   ├── ideogram.go
│   │   └── ideogram_test.go
│   ├── repository/
│   │   └── postgres/
│   │       └── ideogram_repository.go
│   └── usecase/
│       ├── translate_ideogram.go
│       └── translate_ideogram_test.go
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── go.sum
```


