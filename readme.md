**Deployment**

**O Que Faz**:
- **Serviços**: dois serviços Go — `servicoA` (porta `:8080`) e `servicoB` (porta `:8181`).
- **Endpoints**: ambos expõem `POST /temperatura` que recebe JSON `{ "cep": "<cep>" }` e retorna cidade e temperaturas.
- **Observabilidade**: o `otel-collector`, `Zipkin` são definidos no `docker-compose.yaml` para coletar traces/metrics.

**Como Rodar Local**:
- Recomendado (via Docker Compose):

```bash
docker compose up -d --build
```

- Alternativa (rodar serviços Go localmente sem Docker):

```bash
# Em uma aba: servidor A
cd servicoA
go run ./cmd/webserver

# Em outra aba: servidor B
cd ../servicoB
go run ./cmd/webserver
```

- UIs úteis:
- Zipkin - `http://localhost:9411`

- Testar endpoints (exemplos):

Usando o arquivo `api/requestTemp.http` (VSCode REST Client) — já contém exemplos prontos.

Usando `curl`:

```bash
curl -s -X POST http://localhost:8080/temperatura \
  -H "Content-Type: application/json" \
  -d '{"cep":"15771034"}' | jq
```

- Respostas esperadas:
- Sucesso: `200` com JSON `{ "city": "...", "temp_C": <float>, "temp_F": <float>, "temp_K": <float> }`.
- Erro de validação de CEP: `422` com mensagem.
- JSON inválido: `400`.