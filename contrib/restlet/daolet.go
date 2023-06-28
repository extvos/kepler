package restlet

import xql "github.com/archsh/go.xql"

type Identity interface{ string | int32 | int64 }

type Filter struct {
	Operator string
	Values   []interface{}
}

type Order struct {
	Sort  string
	Field string
}

type Query struct {
	Offset  int64
	Limit   int64
	OrderBy []Order
	Filters map[string]Filter
}

type DAO[T interface{ xql.TableIdentified }] struct {
	_table   *xql.Table
	_session *xql.Session
}

func NewDAO[T interface{ xql.TableIdentified }](s *xql.Session, schema ...string) (DAO[T], error) {
	var d DAO[T]
	var t T
	d._table = xql.DeclareTable(t, schema...)
	d._session = s
	return d, nil
}

func (d DAO[T]) SelectById(id interface{}) (*T, error) {
	var t T
	if e := d._session.Table(d._table).Get(id).Scan(&t); nil != e {
		return nil, e
	} else {
		return &t, nil
	}
}

func (d DAO[T]) SelectOne(queries Query) (*T, error) {
	return nil, nil
}

func (d DAO[T]) SelectList(queries Query) ([]T, error) {
	var list []T
	return list, nil
}

func (d DAO[T]) Count(queries Query) (int64, error) {
	return 0, nil
}

func (d DAO[T]) Insert(obj T) (T, error) {
	return obj, nil
}

func (d DAO[T]) Update(obj T, queries Query) (int64, error) {
	return 0, nil
}

func (d DAO[T]) UpdateById(obj T, id interface{}) (int64, error) {
	return 0, nil
}

func (d DAO[T]) Delete(queries Query) (int64, error) {
	return 0, nil
}

func (d DAO[T]) DeleteById(id interface{}) (int64, error) {
	return 0, nil
}
