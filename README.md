# pln-checker

## Setup

### dirnev

- [installation](https://direnv.net/docs/installation.html)
- [setup](https://direnv.net/docs/hook.html)
- create `.envrc` file in the root directory:
    ```
    export PATH=./bin:$PATH
    ```
- run `direnv allow .`

## Usage

The app to perform load testing on host.

Please set environment variable (`LOF_FILE`) to select log file.
You can do it by adding the line from bellow to your `.envrc` file.
```
export LOF_FILE=./logs/log.txt
```
*Remember about refreshing your local environment variables!*
- run `direnv allow .`

`go run  main.go --Host="" --X=10 --Y=1`
- Host - host to test
- X - number of requests per the frequency (Y)
- Y - interval/frequency in seconds