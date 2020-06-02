package fingerprinting

import "testing"

func TestNewGoHash(t *testing.T) {
	type args struct {
		labels map[string]string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "Testing", args: struct{ labels map[string]string }{labels: map[string]string{"testing": "tester"}}, want: 5407378286431217024},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGoHash(tt.args.labels); got != tt.want {
				t.Errorf("NewGoHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result uint64

func BenchmarkNewGoHash(b *testing.B) {
	var r uint64
	for n := 0; n < b.N; n++ {
		r = NewGoHash(map[string]string{"testing": "tester" + string(b.N)})
	}

	result = r
}
