# go-fizzbuzz

FizzBuzz (kind of) REST API in Go.

## Build & Run :construction_worker:

```shell script
make run
```
Check [Makefile](./Makefile) for further commands.

## Documentation :blue_book:

To access Swagger documentation and make some tries, please run project locally and open : http://127.0.0.1:3333/api/0.6.2/docs/index.html

:warning:

When testing request on Swagger, you need to set same sheme for request as scheme on swagger url (http / https)

## Deployment :rocket:
For production deployment, please set CONFIG_PATH environment variable which sets the properties file location.