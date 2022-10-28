package api

import ( 
    "bitbucket.com/jcalvert/or-service/fibre/psql_orm"
)

func Query() *psql_orm.HellOr {
    res := psql_orm.SearchDbHello()
    return res
}
