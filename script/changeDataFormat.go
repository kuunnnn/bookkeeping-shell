package main

import (
	read_data "bookkeeping-shell/read-data"
	"bookkeeping-shell/store"
)

func transformStructureType(r []*read_data.Record) []*store.Record {
	list := make([]*store.Record, 0)
	for _, val := range r {
		list = append(list, &store.Record{
			Type:      val.Type,
			Timestamp: val.Timestamp,
			Desc:      val.Desc,
			Date:      val.Date,
			Money:     val.Money,
			Category:  val.Category,
		})
	}
	return list
}
func main() {
	jsonList, err := read_data.ReadDataToRecordSlice()
	if err != nil {
		panic(err)
	}
	err = store.Rewrite(transformStructureType(jsonList))
	if err != nil {
		panic(err)
	}
}
