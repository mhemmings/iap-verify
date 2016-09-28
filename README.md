#iap-verify

A small CLI tool to verify and get info on individual Apple App Store receipts, using [dogenzaka's go-iap library](https://github.com/dogenzaka/go-iap)

##Installation:

`go install github.com/mhemmings/iap-verify`

##Usage:

`iap-verify /path/to/b64receipt`

To verify a sandbox receipt:

`iap-verify -sandbox /path/to/b64receipt`

To verify an autorenewing subscription, you need a shared secret:

`iap-verify -secret="sharedsecret" /path/to/b64receipt`

##Limitations/Future:

- iOS only, trivial to support Android and Amazon
- Receipts have to be saved to a file, trivial to support stdin
