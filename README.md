# maria-driver

## Install

```
$ go get github.com/innotechdevops/maria-driver
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
    MaxLifetime:  os.Getenv("MARIA_MAX_LIFETIME"),
    MaxIdleConns: os.Getenv("MARIA_MAX_IDLE_CONNS"),
    MaxOpenConns: os.Getenv("MARIA_MAX_OPEN_CONNS"),
})
```