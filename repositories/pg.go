package repositories

const dbname = ""
const host = ""
const port = "5432"

const user = ""
const password = ""
const sslmode = "require"

//os.Getenv("DATABASE_URL")

const PgConnection = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname
