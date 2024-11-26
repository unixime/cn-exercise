# cn-exercise
CodeNotary test exercise


```shell
docker run -p 8081:8080 -p 3322:3322 --rm --name immudb codenotary/immudb:latest
```

```shell
curl 127.0.0.1:8080/foo -d '{"name": "john", "bank_account": 1234, "address":"fdkjlsdjf", "amount": 1.2345, "type": 1}'
```