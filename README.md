# maria-driver

## Install

```
$ go get github.com/innotechdev/maria-driver
```

## How to use

- Wtih env

```golang
driver := mariadriver.New(mariadriver.ConfigEnv())
```

- With config

```golang
driver := mariadriver.New(mariadriver.Config{
    User:         os.Getenv("MARIA_USER"),
    Pass:         os.Getenv("MARIA_PASS"),
    Host:         os.Getenv("MARIA_HOST"),
    DatabaseName: os.Getenv("MARIA_DATABASE"),
    Port:         mariadriver.DefaultPort,
})
```