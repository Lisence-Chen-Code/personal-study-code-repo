package design_modes

import "sync"

//sqlSessionFactory, table, where, like, in, hook

type TableBaseOpt interface {
	Before() *SqlSessionFac
	Where(sql interface{}, args ...interface{}) *SqlSessionFac
	Like(sql interface{}, args ...interface{}) *SqlSessionFac
	In(sql interface{}, args ...interface{}) *SqlSessionFac
	After() *SqlSessionFac
	Table(tableName string) *SqlSessionFac
}

type SqlSessionFac struct {
	RawSql string
}

//单例sql工厂
var once sync.Once = sync.Once{}
var sqlSessionFactory *SqlSessionFac

func NewSqlInstance() *SqlSessionFac {
	once.Do(func() {
		sqlSessionFactory = new(SqlSessionFac)
	})
	return sqlSessionFactory
}

//implement intf
