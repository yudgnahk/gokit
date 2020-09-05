package templates

const RepositoryTemplate = `package repositories

import (
	"MODULE_NAME/adapters"
)

// REPOSITORY_NAMERepository interface
type REPOSITORY_NAMERepository interface {

}

type repository_nameRepository struct {
	orm database.DBAdapter
}

// NewREPOSITORY_NAMERepository return new REPOSITORY_NAME Repository
func NewREPOSITORY_NAMERepository(orm database.DBAdapter) REPOSITORY_NAMERepository {
	return &repository_nameRepository{
		orm: orm,
	}
}

// TODO: Implement DoSomething
func (r *repository_nameRepository) DoSomething() error {
	return nil
}
`
