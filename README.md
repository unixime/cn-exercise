# Home Assignment - Fullstack/Golang Engineer position

Please have a look at our cloud service `immudb Vault` and build a simple application (Backend + Frontend) around it with the following requirements:

* Application is storing accounting information within immudb Vault with the following structure: 
  * account number (unique)
  * account name
  * iban
  * address
  * amount
  * type (sending, receiving)
* Application has an API to add and retrieve accounting information
* Application has a frontend that displays accounting information and allows to create new records. ( **NOT IMPLEMENTED** )

The solution should:
* Have a readme
* Have a documented API
* Have docker-compose so it is easy to run.

Resources:

* immudb Vault documentation: https://vault.immudb.io/docs/
* API reference: https://vault.immudb.io/docs/api/v1


## Missing requirements:

* application frontend

## ToDo

* unit tests
  * add `immodb vault` mock
* code factorization :
  * minimize code replication



