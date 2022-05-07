package branch

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGenerateName(t *testing.T) {
	type args struct {
		num   int
		title string
	}
	for _, tt := range []struct {
		args

		name    string
		wantErr bool
		want    string
	}{
		{
			name: "negative index",
			args: args{
				num:   -1,
				title: "a-title",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "zero index",
			args: args{
				num:   0,
				title: "a-title",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "empty title",
			args: args{
				num:   33,
				title: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "blank title",
			args: args{
				num:   33,
				title: "   ",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "dot title",
			args: args{
				num:   33,
				title: ".",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "branch name should not end with replacement if title already ends with it",
			args: args{
				num:   1,
				title: "-",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "branch name should not end with replacement if title already ends with it",
			args: args{
				num:   1,
				title: "something-ending-with-hyphen-",
			},
			want:    "1-something-ending-with-hyphen",
			wantErr: false,
		},
		{
			name: "branch name should not have non-alphanumeric character",
			args: args{
				num:   777,
				title: "[Feature request] an awesome feature for `tool`",
			},
			want:    "777-feature-request-an-awesome-feature-for-tool",
			wantErr: false,
		},
		{
			name: "ending dot should be ignored",
			args: args{
				num:   1234,
				title: "Bug when running something.",
			},
			want:    "1234-bug-when-running-something",
			wantErr: false,
		},
		{
			name: "branch name should be all lowercase",
			args: args{
				num:   2345,
				title: "THIS IS AN IMPROVEMENT SUGGESTION :-) !!!",
			},
			want:    "2345-this-is-an-improvement-suggestion",
			wantErr: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateName(tt.num, tt.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("got err: %v, but wanted err set to %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
