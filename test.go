package main

import (
	"fmt"
	"github.com/olongfen/gorm-gin-admin/src/pkg/query"
)

func main() {
	q, _ := query.NewQuery(1, 10).ValidCond(map[string]interface{}{"role$lt$": 2})
	fmt.Println(q)
}
