## Database Connector
lightweight_api will auto create Connector if you set environment variable DB_DRIVER_NAME equal to "mysql" or "sqlite".

If you want to use a sql driver other than mysql or sqlite, you can init Connector by yourself.

like 
```go
import	"github.com/ssst0n3/lightweight_db/example/mysql"

var Conn lightweight_db.Connector

func InitConnector() {
	Conn = mysql.Conn()
	lightweight_api.Conn = Conn
}
```

or 

```go
import 	_ "github.com/go-sql-driver/mysql"

var Conn = lightweight_api.InitConnector("mysql", lightweight_db.GetDsnFromEnvNormal())
```