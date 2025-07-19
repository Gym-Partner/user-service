package repository

import (
	"github.com/Gym-Partner/api-common/errs"
	"github.com/Gym-Partner/api-common/status"
	"github.com/Gym-Partner/user-service/internal/domain"
	"gorm.io/gorm"
	"strings"
)

// Repository provides access to the user data stored in the database.
type Repository struct {
	DB      *gorm.DB
	Catalog *status.StatusCatalog
}

// New creates and returns a new Repository instance
// using the provided GORM database handler and logger
func New(db *gorm.DB, catalog *status.StatusCatalog) *Repository {
	return &Repository{
		DB:      db,
		Catalog: catalog,
	}
}

// IsExist implements IRepository.IsExist
func (r *Repository) IsExist(data, OPT string) bool {
	var user domain.User
	var queryColumn string

	switch strings.ToUpper(OPT) {
	case "ID":
		queryColumn = "id"
	case "EMAIL":
		queryColumn = "email"
	}

	if raw := r.DB.
		Where(queryColumn+" = ?", data).
		First(&user); raw.Error != nil {
		r.Catalog.Log("", raw.Error.Error())
		return false
	}

	if user.ID == "" {
		r.Catalog.Log("")
		return false
	} else {
		return true
	}
}

// Create implements IRepository.Create
func (r *Repository) Create(data domain.User) (domain.User, *errs.Error) {
	if raw := r.DB.
		Create(&data); raw.Error != nil {
		r.Catalog.Log("", raw.Error.Error())
		return domain.User{}, errs.New(r.Catalog, "", nil, raw.Error.Error())
	}
	return data, nil
}

// GetAll implements IRepository.GetAll
func (r *Repository) GetAll() (domain.Users, *errs.Error) {
	var users domain.Users

	if raw := r.DB.
		First(&users); raw.Error != nil {
		r.Catalog.Log("", raw.Error.Error())
		return domain.Users{}, errs.New(r.Catalog, "", nil, raw.Error.Error())
	}
	return users, nil
}

// GetOneByID implements IRepository.GetOneByID
func (r *Repository) GetOneByID(uid string) (domain.User, *errs.Error) {
	var user domain.User

	if raw := r.DB.
		Where("id = ?", uid).
		First(&user); raw.Error != nil {
		r.Catalog.Log("", raw.Error.Error())
		return domain.User{}, errs.New(r.Catalog, "", nil, raw.Error.Error())
	}
	return user, nil
}

// GetOneByEmail implements IRepository.GetOneByEmail
func (r *Repository) GetOneByEmail(email string) (domain.User, *errs.Error) {
	var user domain.User

	if raw := r.DB.
		Where("email = ?", email).
		First(&user); raw.Error != nil {
		r.Catalog.Log("", raw.Error.Error())
		return domain.User{}, errs.New(r.Catalog, "", nil, raw.Error.Error())
	}
	return user, nil
}
