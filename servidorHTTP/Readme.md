# Servidor HTTP com GO - Agendamento de Consultas Médicas

## Visão Geral
Este projeto é um servidor HTTP desenvolvido em GoLang adaptado para o contexto de **agendamento de consultas médicas**. Ele permite a criação (agendar), visualização, atualização (reagendar) e exclusão (cancelar) de consultas médicas. A aplicação utiliza PostgreSQL como banco de dados relacional e fornece uma interface web simples e limpa para interação dos pacientes.

---

## Estrutura do Projeto
### Diretórios Principais:
- **`app/`**: Contém a lógica do servidor, incluindo handlers e utilitários.
  - **`handlers/`**: Lida com as requisições HTTP e processa os dados de agendamento.
  - **`utils/`**: Contém funções auxiliares, como conexão ao banco de dados e as queries SQL (CRUD).
- **`static/`**: Contém os arquivos estáticos (HTML e CSS) que compõem o front-end da aplicação.
  - **`forms/`**: Formulários HTML para agendar, visualizar, reagendar e cancelar consultas.
  - **`styles/`**: Arquivos CSS planos para estilização das páginas.

---

## Configuração do Ambiente

### Pré-requisitos
1. **GoLang**: Certifique-se de que o Go está instalado na sua máquina.
2. **PostgreSQL**: Banco de dados utilizado pelo projeto.
3. **Docker** (opcional): Para facilitar a configuração do banco de dados.

### Passos para Configuração
1. **Clone o repositório**:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd servidorHTTP
   ```

2. **Configure o arquivo `.env`**:
   Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```env
   DB_USER=<seu_usuario>
   DB_PASSWORD=<sua_senha>
   DB_NAME=<nome_do_banco>
   DB_HOST=<host_do_banco>
   DB_PORT=<porta_do_banco>
   ```

3. **Configuração do Banco de Dados**:
   - Se estiver usando Docker, utilize o arquivo `docker-compose.yml` para subir o banco de dados:
     ```bash
     docker compose up -d
     ```
   - Crie a tabela de consultas médicas no banco de dados com o seguinte comando SQL:
     ```sql
     CREATE TABLE consultas (
         id SERIAL PRIMARY KEY,
         nome_paciente VARCHAR(150) NOT NULL,
         nome_medico VARCHAR(150) NOT NULL,
         especialidade VARCHAR(150) NOT NULL,
         data_consulta DATE NOT NULL,
         horario_consulta TIME NOT NULL,
         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
     );
     ```

4. **Instale as dependências**:
   Execute o comando abaixo para instalar as dependências do projeto:
   ```bash
   go mod tidy
   ```

---

## Executando o Projeto
1. **Inicie o servidor**:
   ```bash
   go run app/main.go
   ```

2. **Acesse a aplicação**:
   O servidor estará disponível no endereço exibido no terminal, geralmente algo como:
   ```
   http://127.0.0.1:3000/
   ```

---

## Funcionalidades
### Rotas Principais:
- **`/`**: Página inicial (`index.html`) com links para os formulários de agendamento.
- **`/forms/agendarConsulta.html`**: Formulário para agendar uma nova consulta.
- **`/forms/visualizarConsultas`**: Página de listagem que exibe todas as consultas cadastradas em uma tabela padronizada e limpa.
- **`/forms/reagendarConsulta.html`**: Formulário para reagendar (atualizar) dados de uma consulta existente.
- **`/forms/cancelarConsulta.html`**: Formulário para cancelar (excluir) uma consulta agendada por meio do código ID.

### Handlers:
- **`CreateConsultaHandler`**: Processa a criação e inserção de novos agendamentos no banco de dados.
- **`ListConsultaHandler`**: Consulta todas as linhas da tabela no PostgreSQL e injeta os dados no template HTML de forma dinâmica.
- **`UpdateConsultaHandler`**: Atualiza as informações da consulta selecionada de forma parcial ou integral.
- **`DeleteConsultaHandler`**: Remove permanentemente o registro da consulta a partir de seu ID.

---

## Estrutura de Pastas
```
.env
docker-compose.yml
go.mod
app/
  main.go
  handlers/
    createConsultaHandler.go
    listConsultaHandler.go
    updateConsultaHandler.go
    deleteConsultaHandler.go
  utils/
    connectToDB.go
    createConsulta.go
    getConsultas.go
    updateConsulta.go
    deleteConsulta.go
static/
  index.html
  forms/
    agendarConsulta.html
    visualizarConsultas.html
    reagendarConsulta.html
    cancelarConsulta.html
  styles/
    index.style.css
    agendarConsulta.style.css
    visualizarConsultas.style.css
    reagendarConsulta.style.css
    cancelarConsulta.style.css
```

---

## Observações
- Certifique-se de que o banco de dados está configurado corretamente e a tabela `consultas` foi criada antes de iniciar o servidor.
- O projeto utiliza o driver `github.com/lib/pq` para conexão com o PostgreSQL.
- Para mais informações, consulte os comentários no código ou entre em contato com o desenvolvedor.
