package smartbatching

const (
	addOp int = iota
	finishOp
)

type opItem struct {
	kind int
	key  string
	data itemBatch
}
type itemBatch struct {
	rep  chan interface{}
	data interface{}
}

type processBatch interface {
	Do(key string, datas []interface{}) []interface{}
}
type SmartBatching struct {
	p      processBatch
	chanOp chan opItem
}

func (s *SmartBatching) run() {
	batchMap := make(map[string][]itemBatch)
	for op := range s.chanOp {
		key := op.key
		batch, exist := batchMap[key]
		if op.kind == addOp {
			if exist { // key on processing
				batchMap[key] = append(batchMap[key], op.data)
				batch = nil

			} else { // mark key on processing with batch
				batchMap[key] = nil
				batch = []itemBatch{op.data}
			}

		} else if op.kind == finishOp {
			if len(batch) == 0 { // no item remain
				delete(batchMap, key)

			} else { // mark key on processing with batch
				batchMap[key] = nil
			}

		}
		if len(batch) > 0 {

			go func() {
				datas := make([]interface{}, 0, len(batch))
				for _, item := range batch {
					datas = append(datas, item.data)
				}
				for i, rep := range s.p.Do(key, datas) {
					batch[i].rep <- rep
				}
				s.chanOp <- opItem{
					kind: finishOp,
					key:  key,
				}
			}()

		}
	}

}
func (s *SmartBatching) Add(key string, data interface{}) interface{} {
	rep := make(chan interface{})
	s.chanOp <- opItem{addOp, key, itemBatch{rep, data}}
	return <-rep

}
func NewSmartBatching(p processBatch) *SmartBatching {
	s := &SmartBatching{
		p:      p,
		chanOp: make(chan opItem),
	}
	go s.run()
	return s
}
