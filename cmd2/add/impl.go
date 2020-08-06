package add

import (
	read_data "bookkeeping-shell/read-data"
)

// 新增一条记录
func add(money float64, kind string, desc string, day int, typ read_data.Species) error {
	d, t := read_data.GetOffsetTime(day)
	return read_data.Append(&read_data.Record{
		Date:d,
		Timestamp:t,
		Category:kind,
		Type: string(typ),
		Desc:desc,
	})
}
