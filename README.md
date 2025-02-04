# REST-API

REST API для финансовых операций с тремя ручками: 
1) Пополнение баланса пользователя
```bash
curl -X POST http://localhost:8080/deposit -H "Content-Type: application/json" -d '{"userid": 1, "amount": 100}'
```
2) Перевод денег от одного пользователя к другому
```bash
curl -X POST http://localhost:8080/transfer \
-H "Content-Type: application/json" \
-d '{"fromuserid": 1, "touserid": 2, "amount": 50}'

```

3) Просмотр последних операций пользователя
```bash
curl -X GET http://localhost:8080/transactions/1
```

