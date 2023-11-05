package calculations

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		number  int64
		logging bool
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"valid data test 1", args{12, false}, 479001600, false},
		{"valid data test 2", args{0, false}, 1, false},
		{"invalid data test 1", args{-1, false}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateFactorial(tt.args.number, tt.args.logging)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
