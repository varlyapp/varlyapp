package lib

import "testing"

func TestGenerateDNA(t *testing.T) {
	tests := []struct {
		name string
		layers []string
		want string
	}{
		{
			name: "can hash layers",
			layers: []string{"path/to/asset1/.png", "path/to/asset2.png", "path/to/asset3.png"},
			want: "deb38c0c5afbdc5e1dd5d2a063de2c9d34bfcd44",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateDNA(tt.layers); got != tt.want {
				t.Errorf("GenerateDNA() = %v, want %v", got, tt.want)
			}
		})
	}
}
