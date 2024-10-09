# playground-go-sqlc

### Pre-requisites

- go
- relational database (mysql or postgres)
- sqlc

### Generate

```
sqlc generate
```

> [!NOTE]
> sqlc will generate a `db` package from source `sqlc.yaml`

### Run Application

```
go run main.go
```

### Reference

- <https://go.dev/doc/install>
- <https://docs.sqlc.dev/en/latest/overview/install.html>
- <https://medium.com/gravel-engineering/using-sqlc-for-orm-alternative-in-golang-ft-go-migrate-pgx-b9e35ec623b2>
- <https://medium.com/gravel-engineering/using-custom-struct-for-jsonb-type-in-sqlc-f0350d4df967>
