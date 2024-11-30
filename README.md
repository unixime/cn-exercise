# cn-exercise
CodeNotary test exercise


```shell
docker run -p 8081:8080 -p 3322:3322 --rm --name immudb codenotary/immudb:latest
```

```shell
curl 127.0.0.1:8080/foo -d '{"name": "john", "bank_account": 1234, "address":"fdkjlsdjf", "amount": 1.2345, "type": 1}'
```

```shell
curl -X 'DELETE' 'https://vault.immudb.io/ics/api/v1/ledger/default/collection/transactions' -H 'accept: application/json'   -H 'X-API-Key: default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y'   -H 'Content-Type: application/json'



curl -X 'PUT' 'https://vault.immudb.io/ics/api/v1/ledger/default/collection/transactions' -H 'accept: application/json'   -H 'X-API-Key: default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y'   -H 'Content-Type: application/json' -d '
{
  "fields": [
  {
    "name": "uuid",
    "type": "STRING"
  },
  {
    "name": "name",
    "type": "STRING"
  },
  {
    "name": "iban",
    "type": "STRING"
  },
  {
    "name": "address",
    "type": "STRING"
  },
  {
    "name": "amount",
    "type": "DOUBLE"
  },
  {
    "name": "type",
    "type": "INTEGER"
  }
  ],
  "indexes": [
  {
    "fields": [
      "uuid"
    ],
    "isUnique": false
  },
  {
    "fields": [
      "name"
    ],
    "isUnique": false
  },
  {
    "fields": [
      "name",
      "type"
    ],
    "isUnique": false
  }
  ]
}'


curl -X 'PUT' 'https://vault.immudb.io/ics/api/v1/ledger/default/collection/transactions/document' -H 'accept: application/json'   -H 'X-API-Key: default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y'   -H 'Content-Type: application/json' -d '
{
  "uuid" : "100",
  "name" : "John Blake",
  "iban" : "IT32C0300203280141759145451",
  "address" : "foo",
  "amount" : 50,
  "type" : 1
}'

```