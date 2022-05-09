package repositories

const dbname = ""
const host = ""
const port = ""
const user = ""
const password = ""
const sslmode = "require"

//os.Getenv("DATABASE_URL")

const PgConnection = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname

func New() *Queries {
	return &Queries{}
}

type Queries struct {
}
