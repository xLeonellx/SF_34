package calculate

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		inputFile  string
		outputFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name:    "ok",
			args:    args{inputFile: "input_test.txt", outputFile: "output_test.txt"},
			wantErr: false,
		},
		{
			name:    "without file",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Calculate(tt.args.inputFile, tt.args.outputFile); (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
