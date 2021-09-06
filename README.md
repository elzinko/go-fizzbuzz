# go-fizzbuzz

FizzBuzz (kind of) REST API in Go.

## :construction_worker: Build & Run

```shell script
make run
```
Check [Makefile](./Makefile) for further commands.

## :blue_book: Documentation

To access Swagger documentation and make some tries, please run project locally and open : http://127.0.0.1:3333/api/0.6.2/docs/index.html

:warning: When testing request using Swagger, you need to set same sheme for request as the one set on swagger url (http / https)

## :rocket: Deployment

For production deployment, please set CONFIG_PATH environment variable which sets the properties file location.

## :next_track_button: Missing features

- Versionning handling in code
- graceful shutdown
- multiple errors handling
- use gin.TestMode for true Gin unit test
- ci/cd : 
    - better tag handling
    - image push on registry
    - k8s deployment sample (GCP free tier)
    - continuous delivery