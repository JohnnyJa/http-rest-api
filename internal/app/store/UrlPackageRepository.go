package store

import "github.com/JohnnyJa/http-rest-api/internal/app/model"

type UrlPackageRepository struct {
	store *Store
}

func (r *UrlPackageRepository) Create(u *model.UrlPackage) (*model.UrlPackage, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO UrlPackage (Url) OUTPUT Inserted.ID VALUES (@p1)", 
		u.UrlString,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UrlPackageRepository) FindById(id int) (*model.UrlPackage, error) {

	u := &model.UrlPackage{}
	if err := r.store.db.QueryRow(
		"SELECT id, url FROM UrlPackage Where id = @p1", 
		id,
	).Scan(
		&u.ID, 
		&u.UrlString,
	); err !=nil {
		return nil, err
	}
	return u, nil
}
