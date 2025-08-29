package goenv

import "testing"

func TestGet(t *testing.T) {
	type args struct {
		key string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "valid GOVERSION",
			args:    args{key: "GOVERSION"},
			want:    "go1.25.0",
			wantErr: false,
		},
		{
			name:    "invalid key returns error",
			args:    args{key: "__THIS_GO_ENV_KEY_DOES_NOT_EXIST__"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
