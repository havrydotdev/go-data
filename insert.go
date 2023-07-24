package data

import "reflect"

func (r *SQLRepository[K, T]) Insert(entity T) (K, error) {
	var id K
	v := reflect.ValueOf(entity)
	values := []any{}
	for _, c := range r.columns {
		f := v.FieldByName(c)
		values = append(values, f.Interface())
	}
	println("Insert query: ", r.qs.insert)
	row := r.db.QueryRow(r.qs.insert, values...)
	if err := row.Err(); err != nil {
		return id, err
	}

	err := row.Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
