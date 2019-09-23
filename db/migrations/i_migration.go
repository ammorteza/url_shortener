package migrations

type MigrationInterface interface {
	Make() error
	Drop() error
	Reset() error
}
