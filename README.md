# Cryptopals Solutions

My solutions as I slowly work my way through the Cryptopal exercises with the 
other bit twiddlers in the AppSec team at `$WORK`.

Work in progress (until I declare defeat atleast).

## Install 

To pull it down and build it:
```
go get -v github.com/werrett/cryptopals
```

To use run challenges with input:
```
cryptopals
cryptopals ch1 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
```

To test solution based on example input:
```
$GOPATH/src/github.com/werrett/cryptopals
go test ./... -v
```

## Author 

Jonathan Werrett <jonathan@werrett.co>, and the other bit twiddlers at `$WORK`.