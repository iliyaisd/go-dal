package dal

// ISchema represents DAL schema methods
type ISchema interface {
	Select(tableName string) IQuery
	Update(tableName string) IQuery
	Delete(tableName string) IQuery
	Insert(tableName string) IQuery
	AddTable(name string, fields []string) error
	Table(name string) (t *Table)
}
