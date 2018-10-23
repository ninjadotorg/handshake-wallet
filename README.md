# Requirements
### Glide
https://github.com/Masterminds/glide

### Go
https://golang.org/doc/install

# Setup
```bash
go get github.com/ninjadotorg/handshake-wallet
cd $GOPATH/src/github.com/ninjadotorg/handshake-wallet
glide install
```

# Configure
```bash
cd $GOPATH/src/github.com/ninjadotorg/handshake-wallet
cp config/conf.yaml.default config/conf.yaml
```

Edit `config/conf.yaml` to fix your config

# Migrate db
Mysql. create database if not exists

`CREATE DATABASE database CHARACTER SET utf8 COLLATE utf8_general_ci;`

```bash
go run migrate.go
```

# Run server
```bash
go run main.go
```
