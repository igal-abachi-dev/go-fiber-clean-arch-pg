package repositories

const dbname = "d66kr61430h4e"
const host = "ec2-34-194-73-236.compute-1.amazonaws.com"
const port = "5432"

const user = "llbavbgsduhpsz"
const password = "8698120fda348c13ecb5b051ddbddd9fc3a5fb7016895eb3e77c5747836079cd"
const sslmode = "require"

//os.Getenv("DATABASE_URL")

const PgConnection = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname
