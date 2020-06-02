package fingerprinting

import "testing"

func TestNewPromHash(t *testing.T) {
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
			if got := NewPromHash(tt.args.labels); got != tt.want {
				t.Errorf("NewPromHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewPromHash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewPromHash(map[string]string{"testing": "tester" + string(b.N)})
	}
}
