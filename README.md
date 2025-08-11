# Financial API - Uma Jornada DevSecOps

![alt text](https://img.shields.io/docker/pulls/nickolss/financial_api.svg)

## Sobre o Projeto
Esta é uma API REST para gerenciamento de finanças pessoais, desenvolvida em Go. No entanto, seu principal objetivo é servir como um campo de provas para o aprendizado e aplicação de práticas de DevSecOps.
O projeto demonstra como integrar a segurança em cada etapa do ciclo de vida de desenvolvimento de software (Shift-Left Security), utilizando uma pipeline automatizada com ferramentas open-source para garantir a qualidade, a segurança e a integridade do código e dos artefatos gerados.

## Funcionalidades

A API oferece funcionalidades básicas para controle de despesas:

*   **Gestão de Usuários e Autenticação:**
    *   Registro e login com token JWT.
    *   Consulta de dados de usuários.
*   **Operações Financeiras:**
    *   CRUD completo para despesas, receitas e transferências.
*   **Gestão de Categorias:**
    *   Criação e listagem de categorias para classificar as operações.
*   **Relatórios (Futuro):**
    *   Endpoints planejados para resumos financeiros mensais e por categoria.

## Tecnologias Utilizadas
| Ferramenta          | Descrição                                         |
| :------------------ | :------------------------------------------------ |
| **Backend**         | **Go**                                            |
| **Banco de Dados**  | **PostgreSQL**                                    |
| **Containerização** | **Docker & Docker Compose**                       |
| **CI/CD & DevSecOps** | **GitHub Actions, Gitleaks, Trivy** |

## A Pipeline DevSecOps

O coração deste projeto é a pipeline de Integração e Entrega Contínua (CI/CD) configurada com GitHub Actions. Ela automatiza a verificação de segurança em cada commit ou Pull Request.

**Fluxo da Pipeline (em implementação):**

`Commit/PR` → `GitHub Actions` → `[Gitleaks Scan]` → `[Docker Build]` → `[Trivy Scan]` → `Publica Imagem`

#### Ferramentas Integradas:

1.  **Gitleaks:**
    *   ***O que faz?*** Realiza uma varredura no repositório em busca de segredos codificados (senhas, chaves de API, tokens). A pipeline falhará se qualquer segredo for encontrado.

2.  **Trivy:**
    *   ***O que faz?*** Escaneia a imagem Docker gerada em busca de vulnerabilidades conhecidas (CVEs) no sistema operacional base e nas dependências da aplicação.