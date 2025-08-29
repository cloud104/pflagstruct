package goenv

import (
	path "path"
	"testing"
)

func TestGoListDir(t *testing.T) {
	goroot, err := Get("GOROOT")
	if err != nil {
		t.Fatalf("failed to retrieve GOROOT using go env: %v", err)
	}

	type args struct {
		pkg string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "stdlib package fmt",
			args:    args{pkg: "fmt"},
			want:    path.Join(goroot, "src/fmt"),
			wantErr: false,
		},
		{
			name:    "nonexistent package returns error",
			args:    args{pkg: "not/a/real/package/definitelynot"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GoListDir(tt.args.pkg)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoListDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("GoListDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}
