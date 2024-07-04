# Distributed Cache in Go

Scalable caching layer for a distributed system. This project implements a high-performance, distributed caching layer in Go, designed for scalability and consistency. 

## Features
- High performance
- Horizontal scalability
- Data consistency
- Multi-threading support
- Fault tolerance
- Monitoring and logging

## Installation 

```sh
git clone https://github.com/aqwervinker/GO-distributed-cache.git
cd GO-distributed-cache
go mod tidy
```
## Usage

To run the distributed cache server:
```sh
go run main.go
```
## Configuration

Make sure to configure your PostgreSQL connection in db/db.go.

## Monitoring

Metrics are exposed at /metrics endpoint and can be scraped by Prometheus.

```sh
#### `go.mod`

```go
module distributed-cache-go

go 1.20

require (
    github.com/gin-gonic/gin v1.8.1
    github.com/prometheus/client_golang v1.11.0
    github.com/stretchr/testify v1.7.0
)
```

# Распределенный кэш в Go

Масштабируемый кэширующий слой для распределенной системы. В этом проекте реализован высокопроизводительный уровень распределенного кэширования в Go, разработанный для обеспечения масштабируемости и согласованности. 

## Особенности
- Высокая производительность
- Горизонтальная масштабируемость
- Согласованность данных
- Поддержка многопоточности
- Отказоустойчивость
- Мониторинг и ведение журнала

## Установка

```sh
git clone https://github.com/aqwervinker/GO-distributed-cache.git
cd GO-distributed-cache
go mod tidy
```
## Использование

Запусти сервер распределенного кэша:
```sh
go run main.go
```
## Настройка

Проверь, что настроено подключение к PostgreSQL в ```db/db.go```

## Мониторинг

Метрики лежат в ```/metrics endpoint``` и могут быть извлечены с помощью Prometheus.

```sh
#### `go.mod`

```go
module go-distributed-cache

go 1.20

require (
    github.com/gin-gonic/gin v1.8.1
    github.com/prometheus/client_golang v1.11.0
    github.com/stretchr/testify v1.7.0
)
```
