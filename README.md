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
├── cmd/          
│   ├── server/
│   │   ├── main.go
├── internal/          # Código-fonte do Backend
│   ├── database/
│   │   ├── db.go
│   ├── entity/
│   │   ├── ideogram_entity.go
│   ├── handlers/
│   │   ├── ideogram_handler.go
│   ├── repository/
│   │   ├── ideogram_repository.go
│   ├── usecase/
│   │   ├── translate_ideogram.go
└── README.md
