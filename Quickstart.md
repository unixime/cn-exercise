# CN-Exercise QuickStart


## Build image

Clone repository
$ git@github.com.piccio:unixime/cn-exercise.git
$ cd cn-exercise
```shell

```
Build the docker image

```shell
$ docker build  -t piccio/cn-exercise:0.0.0 -f build/docker/Dockerfile .
[ ... ]
 => => naming to docker.io/piccio/cn-exercise:0.0.0
```

## Run container

Application needs 2 env variable to run:
* `CN_URL` ( default `https://https://vault.immudb.io`). the `immudb vault` service URL.
* `CN_API_KEY`: the READ/WRITE Api-Key.

```shell
docker run --env CN_API_KEY=<your_private_api_key> -p 8080:8080 piccio/cn-exercise:0.0.0

```

## Examples

Register a new transaction:

```shell
$ curl -X POST http://127.0.0.1:8080/transaction -d '
{
  "address": "Rome",
  "amount": 200,
  "iban": "IT32C0300203280141759145452",
  "name": "Elon Mask",
  "type": 0,
  "uuid": "103"
}'

```

Get transactions based on account `uuid`

```shell
$ curl -X GET http://127.0.0.1:8080/transactions?name=Elon%20Mask
{
    "revisions": [
        {
            "document": {
                "uuid": "103",
                "name": "Elon Mask",
                "iban": "IT32C0300203280141759145452",
                "address": "Rome",
                "amount": 200,
                "type": 0
            }
        }
    ],
    "page": 1,
    "perPage": 100
```

Get transactions based on account `uuid`

```shell
$ curl -X GET http://127.0.0.1:8080/transactions?uuid=103
{
    "revisions": [
        {
            "document": {
                "uuid": "103",
                "name": "Elon Mask",
                "iban": "IT32C0300203280141759145452",
                "address": "Rome",
                "amount": 200,
                "type": 0
            }
        }
    ],
    "page": 1,
    "perPage": 100
```