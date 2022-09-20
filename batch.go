package smartbatching

import "sync"

type item_batch struct {
	rep  chan interface{}
	data interface{}
}
type SmartBatching struct {
	muBatch  *sync.Mutex
	tblBatch map[string][]item_batch
}
type processBatch interface {
	Do(key string, datas []interface{}) []interface{}
}

func (s *SmartBatching) doBatch(p processBatch, key string, items ...item_batch) {
	datas := make([]interface{}, 0, len(items))
	for _, item := range items {
		datas = append(datas, item.data)
	}
	for i, rep := range p.Do(key, datas) {
		items[i].rep <- rep
	}
	s.muBatch.Lock()
	items = s.tblBatch[key]
	if len(items) == 0 {
		delete(s.tblBatch, key)
	} else {
		s.tblBatch[key] = nil
	}
	s.muBatch.Unlock()
	if len(items) > 0 {
		s.doBatch(p, key, items...)

	}

}
func (s *SmartBatching) Add(p processBatch, key string, data interface{}) interface{} {
	rep := make(chan interface{})
	s.muBatch.Lock()
	batch, exist := s.tblBatch[key]
	if exist {
		batch = append(batch, item_batch{rep, data})

	}
	s.tblBatch[key] = batch
	s.muBatch.Unlock()
	if !exist {
		go s.doBatch(p, key, item_batch{rep, data})
	}
	return <-rep

}
func NewSmartBatching() *SmartBatching {
	return &SmartBatching{
		muBatch:  &sync.Mutex{},
		tblBatch: map[string][]item_batch{},
	}
}
