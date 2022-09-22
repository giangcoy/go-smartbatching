package smartbatching

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

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
	s := NewSmartBatching(&processBench{mtx: sync.Mutex{}})
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Add("A", int(100))
		}
	})

}

type processTest struct{}

func (p *processTest) Do(key string, datas []interface{}) []interface{} {
	time.Sleep(100 * time.Millisecond)
	return datas

}
func TestSmartBatching_Add(t *testing.T) {
	s := NewSmartBatching(&processTest{})

	type args struct {
		key  string
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Case add 10 return 10",
			args: args{"A", int(10)},
			want: int(10),
		},
		{
			name: "Case add 20 return 20",
			args: args{"A", int(10)},
			want: int(10),
		},
		{
			name: "Case add 10 return 10",
			args: args{"A", int(10)},
			want: int(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := s.Add(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SmartBatching.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
