package repository

type Repository struct {
	dataMap map[int]struct{}
}

func New() *Repository {
	return &Repository{
		dataMap: make(map[int]struct{}),
	}
}

func (r *Repository) Save(chislo int) {
	r.dataMap[chislo] = struct{}{}
}
