package migrations

type MigrationInterface interface {
	Make()
	Drop()
	Reset()
}
