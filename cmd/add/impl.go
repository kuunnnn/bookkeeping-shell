package add

import (
	"bookkeeping-shell/store"
)

// 新增一条记录
func add(money float64, kind string, desc string, day int, typ store.Species) error {
	d, t := store.GetOffsetTime(day)
	return store.Append(&store.Record{
		Date:      d,
		Money:     money,
		Timestamp: t,
		Category:  kind,
		Type:      string(typ),
		Desc:      desc,
	})
}
