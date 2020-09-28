remember init conn

like 
```go
var Conn lightweight_db.Connector

func InitConnector() {
	Conn = mysql.Conn()
	lightweight_api.Conn = Conn
}
```

or 

```go
lightweight_api.InitConnector("mysql", lightweight_db.GetDsnFromEnvNormal())
```