package smartbatch

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

type processTest struct {
}

func (p *processTest) Do(key string, datas []interface{}) []interface{} {
	return datas
}
func Test_smartbatch_doBatch(t *testing.T) {
	type fields struct {
		muBatch  *sync.Mutex
		tblBatch map[string][]item_batch
	}
	type args struct {
		p     processBatch
		key   string
		items []item_batch
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Case empty",
			fields: fields{muBatch: &sync.Mutex{}, tblBatch: map[string][]item_batch{}},
			args: args{
				p:     &processTest{},
				key:   "A",
				items: []item_batch{{rep: make(chan interface{}, 1), data: int(0)}},
			},
		},
		{
			name: "Case exist",
			fields: fields{muBatch: &sync.Mutex{}, tblBatch: map[string][]item_batch{
				"A": {{rep: make(chan interface{}, 1), data: int(0)}}},
			},
			args: args{
				p:     &processTest{},
				key:   "A",
				items: []item_batch{{rep: make(chan interface{}, 1), data: int(0)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SmartBatching{
				muBatch:  tt.fields.muBatch,
				tblBatch: tt.fields.tblBatch,
			}
			s.doBatch(tt.args.p, tt.args.key, tt.args.items...)
		})
	}
}

func Test_smartbatch_Add(t *testing.T) {
	type fields struct {
		muBatch  *sync.Mutex
		tblBatch map[string][]item_batch
	}
	type args struct {
		p    processBatch
		key  string
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "Case Empty",
			fields: fields{muBatch: &sync.Mutex{}, tblBatch: map[string][]item_batch{}},
			args: args{
				p:    &processTest{},
				key:  "A",
				data: int(10),
			},
			want: int(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SmartBatching{
				muBatch:  tt.fields.muBatch,
				tblBatch: tt.fields.tblBatch,
			}
			if got := s.Add(tt.args.p, tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smartbatch.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

type processBench struct {
	mtx sync.Mutex
}

func (p *processBench) Do(key string, datas []interface{}) []interface{} {
	p.mtx.Lock()
	time.Sleep(50 * time.Millisecond)
	p.mtx.Unlock()
	return datas
}
func Benchmark_Add(b *testing.B) {
	s := &SmartBatching{muBatch: &sync.Mutex{}, tblBatch: map[string][]item_batch{}}
	p := &processBench{mtx: sync.Mutex{}}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Add(p, "A", int(100))
		}
	})

}
