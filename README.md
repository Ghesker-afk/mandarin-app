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

---

## Componentes

- **handlers/:** Responsável por receber e tratar as requisições externas da aplicação. Os handlers interpretam os dados de entrada, acionam os casos de uso correspondentes e transformam seus resultados em respostas adequadas ao protocolo utilizado, como HTTP.

- **usecases/:** Contém a lógica de aplicação e coordena a execução das operações que atendem às necessidades do sistema. Os casos de uso definem o fluxo das operações, acionando os repositórios e aplicando as regras necessárias para realizar cada funcionalidade, sem depender diretamente de detalhes de transporte ou armazenamento.

- **repositories/:** Define os mecanismos de acesso e persistência dos dados utilizados pela aplicação. Os repositórios abstraem os detalhes de armazenamento e recuperação das informações, permitindo que os casos de uso trabalhem com contratos de acesso a dados sem precisar conhecer os detalhes de sua implementação.

- **entities/:** Contém as estruturas que representam os conceitos e objetos centrais do domínio da aplicação. As entidades definem os dados e, quando aplicável, o comportamento associado a esses conceitos, servindo como modelos compartilhados entre as diferentes camadas do sistema.
