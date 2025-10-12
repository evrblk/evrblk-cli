# Everblack CLI

[![Go](https://github.com/evrblk/evrblk-cli/actions/workflows/go.yml/badge.svg)](https://github.com/evrblk/evrblk-cli/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/evrblk/evrblk-cli)](https://goreportcard.com/report/github.com/evrblk/evrblk-cli)

`evrblk` is an official all-in-one CLI tool for all Everblack services. It can be used for:

* calling API methods on all Everblack services:
  * [Grackle](https://github.com/evrblk/grackle) 
* working with API keys (Alfa, Bravo)

## Installation

```shell
go install github.com/evrblk/evrblk-cli/cmd/evrblk
```

## Calling API methods

```shell
evrblk <SERVICE> <METHOD> --endpoint=<ENDPOINT>
```

* `SERVICE` is a versioned name of a service, i.e. `grackle-preview`
* `METHOD` is a method name, i.e. `list-namespaces`
* `ENDPOINT` is an address of API gateway, i.e. `localhost:8000` 

`evrblk` takes a JSON request via a pipe and prints JSON response from the API call:

```shell
# use echo
echo '{"namespace_name": "my_namespace"}' | evrblk grackle-preview list-semaphores --endpoint=localhost:8000

# or use files
cat acquire_lock_request.json | evrblk grackle-preview acquire-lock --endpoint=localhost:8000
```

By default `evrblk` runs in unauthenticated mode. To use authentication add API keys via env variables:

```shell
export EVRBLK_API_KEY_ID=<KEY_ID>
export EVRBLK_API_SECRET_KEY=<KEY_SECRET>
```

## Working with API keys

* `evrblk authn generate-alfa-key` generates Alfa key. Store the private PEM along with your application and pass to
  Everblack SDK. Take public PEM into the cloud or into your own installation (save as a file named with prefix 
  `key_alfa_`).
* `evrblk authn generate-bravo-key` generates Bravo key. Store the secret along with your application and pass to
  Everblack SDK. Take the same secret into your own installation (save as a file named with prefix `key_bravo_`).

## License

Everblack CLI is released under the [MIT License](https://opensource.org/license/mit).
