package extract_ints_numbers_from_str

import (
	"reflect"
	"testing"
)

func TestExtractNumbers(t *testing.T) {
	type args struct {
		query string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Correct case",
			args: args{
				query: `{"ports": {"dst": [[2049]], "src": [[80]]}, "protocol": "TCP"}`,
			},
			want: []string{"2049", "80"},
		},
		{
			name: "Zero port number",
			args: args{
				query: `{"ports": {"dst": [[2049]], "src": [[0]]}, "protocol": "TCP"}`,
			},
			wantErr: true,
		},
		{
			name: "Ports not found",
			args: args{
				query: `{"ports": {"dst": [[abc]], "src": [[efg]]}, "protocol": "TCP"}`,
			},
			wantErr: true,
		},
		{
			name: "Empty string",
			args: args{
				query: ``,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractNumbers(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractNumbers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
