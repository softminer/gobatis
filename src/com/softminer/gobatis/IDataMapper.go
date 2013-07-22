package gobatis

import (
	//  "database/sql/driver"
	"reflect"
	"container/list"
)

type IDataMapper interface {
	Delete(statement string, object interface{}) (affectedCount int, err error)
	Insert(statement string, object interface{}) (resultObject interface{}, err error)
	Update(statement string, object interface{}) (affectedCount int)
	QueryForObject(statement string, objectStruct reflect.Type,param interface{}) (resultObject interface{}, err error)
	QueryForList(statement string, objectStruct reflect.Type,param interface{})  (resultList list.List,err error)
}
