package data

type Fields map[string]interface{}

func (r *SQLRepository[K, T]) GetAll() ([]T, error) {
	res := []T{}
	err := r.db.Select(&res, getAll(T.GetTableName(*new(T))))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *SQLRepository[K, T]) GetById(id K) (T, error) {
	res := *new(T)
	err := r.db.Get(&res, getById(T.GetTableName(res)), id)
	if err != nil {
		return *new(T), err
	}

	return res, nil
}
