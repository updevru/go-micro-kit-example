# Демо использования Micro Kit

Приложение представляет из себя распределенные key/value storage.

### Функционал

- Запись и получение key/value данных по http и gRPC
- Регистрация в Consul с health checks
- Отправка данных OpenTelemetry в коллектор (логи, трейсы, метрики)
- Распределенное хранилище master-master или master-slave
- Фоновая очистка старых данных

### Генерация кода

```bash
protoc -I proto .\proto\store\*.proto --go_out=./gen/ --go_opt=paths=source_relative --go-grpc_out=./gen/ --go-grpc_opt=paths=source_relative --grpc-gateway_out ./gen --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out ./docs --openapiv2_opt allow_merge=true,merge_file_name=api
```

### Запуск 

```bash
docker compose up
```

Будет запущено два endpoint (http и gRPC), Consul, OpenTelemetry collector, Grafana Stack (Grafana, Tempo, Loki, Mimir).

```
time=2024-05-14T23:56:20.592+03:00 level=INFO msg="http server listening at" address=:8080
time=2024-05-14T23:56:20.593+03:00 level=INFO msg="grpc server listening at" address=[::]:8081
```