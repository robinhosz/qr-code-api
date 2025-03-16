# API de Geração de QR Codes

Esta API permite gerar QR codes a partir de um texto ou URL. Ela oferece duas funcionalidades principais:
1. Gerar um QR code em formato **base64** (JSON).
2. Baixar um QR code como uma imagem **PNG**.

---

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação usada para desenvolver a API.
- **Gin**: Framework web para criar rotas e gerenciar requisições HTTP.
- **go-qrcode**: Biblioteca para gerar QR codes.
---

## Como Usar

### Pré-requisitos

- **Go**: Instale o Go a partir do [site oficial](https://golang.org/dl/).
- **Git**: Instale o Git a partir do [site oficial](https://git-scm.com/).

---

### Passo a Passo

#### 1. Clone o Repositório

Clone o repositório para o seu ambiente local:

```bash
git clone https://github.com/seu-usuario/qr-code-api.git
cd qr-code-api
