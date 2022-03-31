# CryptoBot API wrapper in golang
[![CodeQL](https://github.com/ipsavitsky/cryptobotAPI/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/ipsavitsky/cryptobotAPI/actions/workflows/codeql-analysis.yml)
[![Go](https://github.com/ipsavitsky/cryptobotAPI/actions/workflows/go.yml/badge.svg)](https://github.com/ipsavitsky/cryptobotAPI/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ipsavitsky/cryptobotAPI)](https://goreportcard.com/report/github.com/ipsavitsky/cryptobotAPI)
[![Go Reference](https://pkg.go.dev/badge/github.com/ipsavitsky/cryptobotAPI.svg)](https://pkg.go.dev/github.com/ipsavitsky/cryptobotAPI)
## Installation
```
go get github.com/ipsavitsky/cryptobotAPI
```
## Usage

```go
import (
    crypto "github.com/ipsavitsky/cryptobotAPI"
)

func main() {
    api := crypto.CryptoBotAPI{
        Options: crypto.APIOptions{
            Protocol: "https",
            Host:     "pay.crypt.bot",
        },
        Api_key: "<your key here>",
    }
    resp, err := api.GetMe()
    if err != nil {
        panic(err)
    }
    // do whatever :)
}
```

## Test
To run all tests just run:
```
go test -v
```

## Documentation
API docs [here](https://telegra.ph/Crypto-Pay-API-11-25)


## Contributing
Issues and PRs are welcome
## Contact
 - telegram: [@ilya_savitsky](https://t.me/ilya_savitsky)