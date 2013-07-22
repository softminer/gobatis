package gobatis

import (
 //  "database/sql/driver"
)

type IDataMapper interface {
    Delete(statement string,object interface{}) (affectedCount int,err error)
    Insert(statement string,object interface{}) (resultObject interface{},err error) 
    Update(statement string,object interface{}) (affectedCount int)
 //   QueryForObject(statement string,objectStruct type) (resultObject interface{},err error)
    
} 

