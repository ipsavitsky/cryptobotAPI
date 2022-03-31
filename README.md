# CryptoBot API wrapper in golang

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

## Documentation
Docs for the wrapper are coming, meanwhile you can consult the API docs [here](https://telegra.ph/Crypto-Pay-API-11-25)

## Contributing
Issues and PRs are welcome
## Contact
 - telegram: [@ipsavitsky](https://t.me/ilya_savitsky)