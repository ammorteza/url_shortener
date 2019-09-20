package interfaces

type MigrationInterface interface {
	Make()
	Drop()
}

type MigrationCliInterface interface {
	Make(connection DbConnection)
	Reset(connection DbConnection)
}
