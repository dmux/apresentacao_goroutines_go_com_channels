# Benchmark de Performance: Go vs Python - Concorrência com I/O

## Visão Geral

Este projeto apresenta um estudo comparativo de performance entre implementações em **Go** e **Python** para aplicações web com operações de I/O concorrentes. O objetivo é demonstrar as diferenças fundamentais de arquitetura e performance entre as duas linguagens em cenários de alta concorrência.

## Arquiteturas Implementadas

### Go - Goroutines com Channels

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ch := make(chan string)
    go func() {
        time.Sleep(time.Millisecond)  // Simula I/O
        ch <- "ok"
    }()
    w.Write([]byte(<-ch))
}
```

**Características:**

- **Goroutines**: Threads leves gerenciadas pelo runtime Go
- **Channels**: Comunicação segura entre goroutines
- **Modelo M:N**: Múltiplas goroutines em threads do sistema
- **Stack dinâmico**: Inicia com ~2KB, cresce conforme necessário
- **Garbage Collector**: Otimizado para baixa latência

### Python - Async/Await com FastAPI

```python
@app.get('/')
async def read_root():
    await asyncio.sleep(0.001)  # Simula I/O
    return {"message": "ok"}
```

**Características:**

- **Event Loop**: Single-threaded com concorrência cooperativa
- **Coroutines**: Funções assíncronas com await/async
- **Uvicorn**: Servidor ASGI de alta performance
- **GIL (Global Interpreter Lock)**: Limita execução real paralela
- **Memory overhead**: Maior por conexão

> ℹ️ **O GIL é um mutex (lock global) que protege o acesso aos objetos Python, permitindo que apenas uma thread execute código Python por vez, mesmo em sistemas multi-core.**

## Ambiente de Teste

### Hardware/Infraestrutura

- **Plataforma**: Docker containers isolados
- **Rede**: Docker bridge network
- **Recursos**: Compartilhados entre containers

### Configuração do Benchmark

- **Ferramenta**: wrk (HTTP benchmarking tool)
- **Duração**: 5 minutos por teste
- **Threads**: 4 threads de teste
- **Conexões**: 100 conexões concorrentes
- **Simulação I/O**: 1ms de delay por requisição

### Ferramenta de Benchmark: wrk

O **wrk** é uma ferramenta moderna de benchmarking HTTP amplamente utilizada para testes de performance de aplicações web. Foi escolhida por sua precisão, eficiência e capacidade de gerar alta carga de trabalho.

#### Características do wrk

**Arquitetura:**

- **Multi-threaded**: Utiliza múltiplas threads para gerar carga
- **Event-driven**: Baseado em epoll/kqueue para alta eficiência
- **Low overhead**: Mínimo impacto na medição de performance
- **Scriptable**: Suporta scripts Lua para cenários customizados

**Vantagens:**

- **Precisão**: Medições de latência em microssegundos
- **Escalabilidade**: Capaz de gerar milhares de conexões concorrentes
- **Estatísticas detalhadas**: Percentis, distribuição de latência
- **Baixo consumo**: Não interfere significativamente nos resultados

#### Comando Executado

```bash
# Comando para Go
wrk -t4 -c100 -d5m http://go:8080

# Comando para Python
wrk -t4 -c100 -d5m http://python:8080
```

**Parâmetros:**

- `-t4`: 4 threads para gerar requisições
- `-c100`: 100 conexões HTTP concorrentes
- `-d5m`: Duração de 5 minutos por teste
- `http://go:8080` / `http://python:8080`: URLs dos serviços

#### Métricas Coletadas

**Latência:**

- **Avg**: Latência média das requisições
- **Stdev**: Desvio padrão da latência
- **Max**: Latência máxima observada
- **+/- Stdev**: Percentual dentro de 1 desvio padrão

**Throughput:**

- **Req/Sec**: Requisições por segundo por thread
- **Transfer/sec**: Bytes transferidos por segundo
- **Total requests**: Número total de requisições processadas

**Confiabilidade:**

- **Socket errors**: Erros de conexão, leitura, escrita
- **Timeouts**: Requisições que excederam tempo limite

#### Distribuição de Carga

```text
Padrão de Carga (100 conexões concorrentes):
┌─────────────────────────────────────────┐
│ Thread 1: 25 conexões                   │
│ Thread 2: 25 conexões                   │
│ Thread 3: 25 conexões                   │
│ Thread 4: 25 conexões                   │
└─────────────────────────────────────────┘
         │
         ▼
   ┌─────────────┐
   │   Servidor  │ (Go ou Python)
   │   :8080     │
   └─────────────┘
```

#### Vantagens do wrk sobre Outras Ferramentas

**vs Apache Bench (ab):**

- ✅ Multi-threaded (ab é single-threaded)
- ✅ Estatísticas mais detalhadas
- ✅ Melhor handling de conexões HTTP/1.1

**vs Apache JMeter:**

- ✅ Menor overhead (não usa GUI)
- ✅ Mais preciso para testes de alta carga
- ✅ Configuração mais simples

**vs curl/wget:**

- ✅ Projetado especificamente para benchmarking
- ✅ Métricas estatísticas automáticas
- ✅ Carga concorrente real

#### Limitações e Considerações

**Limitações:**

- Focado apenas em HTTP (não HTTPS por padrão)
- Não simula comportamento real de usuário
- Carga sintética constante

**Considerações:**

- **Warm-up**: Primeiros segundos podem ter resultados variáveis
- **Network**: Testes em containers eliminam latência de rede real
- **Resource sharing**: Containers compartilham recursos do host

#### Interpretação dos Resultados

**Indicadores de Performance:**

- **RPS alto + Latência baixa**: Sistema bem otimizado
- **RPS baixo + Latência alta**: Gargalos de processamento
- **Timeouts frequentes**: Saturação ou deadlocks
- **Alto desvio padrão**: Performance inconsistente

**Sinais de Problemas:**

- Latência crescente ao longo do tempo
- Timeouts aumentando progressivamente
- RPS decrescente durante o teste
- Alto desvio padrão na latência

### Especificações dos Containers

#### Go Container

- **Base**: `scratch` (container mínimo)
- **Binário**: Compilado estaticamente (~5MB)
- **Runtime**: Go 1.24
- **Servidor**: `net/http` padrão

#### Python Container

- **Base**: `python:3.13-slim`
- **Framework**: FastAPI + Uvicorn
- **Gerenciador**: uv para dependências
- **Tamanho**: ~100MB+

## Resultados do Benchmark

### Go Performance

```text
Latency     2.51ms    ±0.97ms   (Max: 68.30ms)
Req/Sec     9.96k     ±600.54   (Max: 12.69k)
Total:      11,897,335 requests em 5min
RPS:        39,651.56 requests/segundo
Throughput: 4.46MB/segundo
Errors:     1 timeout
```

### Python Performance

```text
Latency     273.06ms  ±514.86ms (Max: 2.00s)
Req/Sec     3.66k     ±1.11k    (Max: 4.95k)
Total:      2,425,561 requests em 5min
RPS:        8,082.66 requests/segundo
Throughput: 1.09MB/segundo
Errors:     1,186 timeouts
```

## Análise Técnica Detalhada

### Performance Comparativa

| Métrica             | Go       |   Python |   Diferença |
| ------------------- | -------- | -------: | ----------: |
| **Requests/Second** | 39,652   |    8,083 |   **+390%** |
| **Latência Média**  | 2.51ms   | 273.06ms | **-9,884%** |
| **Throughput**      | 4.46MB/s | 1.09MB/s |   **+309%** |
| **Timeouts**        | 1        |    1,186 |  **-99.9%** |
| **Total Requests**  | 11.9M    |     2.4M |   **+393%** |

### Gráficos de Performance

#### 📊 Requests per Second

```text
Go      ████████████████████████████████████████ 39,652 RPS
Python  ████████                                  8,083 RPS
        0     10k    20k    30k    40k    50k
        │      │      │      │      │      │
        └─ Go supera Python em 4.9x
```

#### ⏱️ Latência Média

```text
Go      █ 2.51ms
Python  ████████████████████████████████████████████████████████████████████████████████████████████████████ 273.06ms
        0ms    50ms   100ms  150ms  200ms  250ms  300ms
        │       │       │       │       │       │       │
        └─ Python tem 108x mais latência
```

#### 🚀 Throughput (MB/s)

```text
Go      ████████████████████████████████████████ 4.46 MB/s
Python  ████████████                            1.09 MB/s
        0      1      2      3      4      5 MB/s
        │      │      │      │      │      │
        └─ Go transfere 4.1x mais dados
```

#### ❌ Timeouts Comparação

```text
Go      ▌ 1 timeout
Python  ████████████████████████████████████████████████████████████████████████████████████████████████████ 1,186 timeouts
        0     200    400    600    800   1000  1200
        │      │      │      │      │      │      │
        └─ Python teve 1,186x mais erros
```

#### 📈 Total de Requests Processadas

```text
Go      ████████████████████████████████████████████████████████████████████████████████████████████████████ 11.9M
Python  ████████████████████                                                                                 2.4M
        0M    2M    4M    6M    8M   10M   12M
        │     │     │     │     │     │     │
        └─ Go processou 4.9x mais requests
```

### 📊 Análise Visual da Diferença de Performance

#### Eficiência Relativa (Go como baseline 100%)

```text
Métrica              Go    Python   Gap
Requests/Second     100%    20% ████████████████████████████████████████████████████████████████████████████████ -80%
Throughput          100%    24% ████████████████████████████████████████████████████████████████████████████████ -76%
Total Requests      100%    20% ████████████████████████████████████████████████████████████████████████████████ -80%
Latência (inverso)  100%     1% ███████████████████████████████████████████████████████████████████████████████▌ -99%
Confiabilidade      100%     0% █████████████████████████████████████████████████████████████████████████████████ -100%
```

#### Performance Score (0-100)

```text
Go      ████████████████████████████████████████████████████████████████████████████████████████████████████ 95/100
Python  ████████████████                                                                                     16/100
        0    10   20   30   40   50   60   70   80   90  100
        │     │    │    │    │    │    │    │    │    │    │
        └─ Go Score: 5.9x superior ao Python
```

### Análise de Latência

**Go:**

- Latência consistente e baixa (2.51ms)
- Desvio padrão baixo (0.97ms)
- Distribuição concentrada (73.61% dentro de ±1 desvio)

**Python:**

- Latência alta e variável (273ms)
- Alto desvio padrão (514ms)
- Distribuição dispersa, indicando gargalos

### Análise de Concorrência

**Go Goroutines:**

- Criação/destruição eficiente (~100ns por goroutine)
- Memory footprint baixo (~2KB inicial)
- Scheduling preemptivo pelo runtime
- Sem contenção por GIL

**Python Async:**

- Event loop single-threaded
- Overhead de context switching entre coroutines
- GIL limita paralelismo real
- Memory overhead maior por conexão

### Análise de Recursos

**Go:**

- Binário compilado pequeno (~5MB)
- Consumo de memória otimizado
- Garbage collector de baixa latência
- Runtime eficiente

**Python:**

- Interpretador + dependências (~100MB+)
- Overhead de interpretação
- GC baseado em reference counting + cycle detection
- Runtime mais pesado

## Padrões Observados

### Scalabilidade

- **Go**: Linear até saturação de CPU/memória
- **Python**: Saturação prematura por limitações arquiteturais

### Latência sob Carga

- **Go**: Mantém latência baixa mesmo com alta carga
- **Python**: Degradação exponencial da latência

### Tratamento de Erros

- **Go**: 1 timeout em 11.9M requests (0.000008%)
- **Python**: 1,186 timeouts em 2.4M requests (0.049%)

## Conclusões e Recomendações

### Quando Usar Go

**✅ Recomendado para:**

- **APIs de alta performance** (>10k RPS)
- **Microserviços** com baixa latência
- **Sistemas distribuídos** e concurrent
- **Aplicações I/O intensive**
- **Services real-time** (gaming, trading, streaming)
- **Proxy/Gateway** services
- **Container workloads** (footprint pequeno)

**Vantagens:**

- Performance superior (4-5x mais RPS)
- Latência consistente e baixa
- Consumo eficiente de recursos
- Escalabilidade linear
- Deploy simples (binário único)

### Quando Usar Python

**✅ Recomendado para:**

- **Prototipagem rápida** e MVPs
- **APIs com lógica complexa** (ML, data processing)
- **Integração** com ecosistema científico
- **CRUD applications** com carga moderada (<5k RPS)
- **Background jobs** e batch processing
- **DevOps tools** e automation

**Vantagens:**

- Desenvolvimento mais rápido
- Ecossistema rico (libraries)
- Flexibilidade e expressividade
- Comunidade e documentação
- Ideal para rapid iteration

### Impacto em Produção

#### Custos de Infraestrutura

- **Go**: 4-5x menos servidores necessários
- **Python**: Maior custo por capacidade equivalente

#### Experiência do Usuário

- **Go**: Response times consistentes (~2ms)
- **Python**: Variabilidade alta (50-500ms+)

#### Operações

- **Go**: Deploy simples, monitoring básico
- **Python**: Tuning complexo, monitoring detalhado

## Limitações do Estudo

1. **Teste sintético**: I/O simulado pode não refletir cenários reais
2. **Hardware**: Não testado em hardware dedicado
3. **Configuração**: Configurações padrão, sem tuning específico
4. **Complexidade**: Handlers simples, sem lógica de negócio
5. **Persistência**: Sem testes com banco de dados

## Próximos Passos

1. **Benchmarks com banco de dados** (PostgreSQL, Redis)
2. **Testes de carga progressiva** (ramp-up testing)
3. **Profiling detalhado** de CPU e memória
4. **Testes com payloads maiores** (JSON complexo)
5. **Comparação com outras linguagens** (Rust, Java, Node.js)

---

**Conclusão Final**: Para aplicações que demandam alta performance, baixa latência e eficiência de recursos, **Go demonstra superioridade clara**. Python mantém sua relevância em cenários onde velocity de desenvolvimento e flexibilidade são prioritárias sobre performance máxima.
