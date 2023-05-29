package branch

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateName(t *testing.T) {
	type args struct {
		id    string
		title string
	}
	for _, tt := range []struct {
		args

		name    string
		wantErr bool
		want    string
	}{
		{
			name: "id empty",
			args: args{
				id:    "",
				title: "a-title",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "id blank",
			args: args{
				id:    "       ",
				title: "a-title",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "id starting with an hyphen -",
			args: args{
				id:    "-1",
				title: "a-title",
			},
			want:    "1-a-title",
			wantErr: false,
		},
		{
			name: "any other id",
			args: args{
				id:    "myId",
				title: "a-title",
			},
			want:    "myid-a-title",
			wantErr: false,
		},
		{
			name: "empty title",
			args: args{
				id:    "33",
				title: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "blank title",
			args: args{
				id:    "33",
				title: "   ",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "dot title",
			args: args{
				id:    "33",
				title: ".",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "branch name should not end with replacement if title already ends with it",
			args: args{
				id:    "1",
				title: "-",
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "branch name should not end with replacement if title already ends with it",
			args: args{
				id:    "1",
				title: "something-ending-with-hyphen-",
			},
			want:    "1-something-ending-with-hyphen",
			wantErr: false,
		},
		{
			name: "branch name should not have non-alphanumeric character",
			args: args{
				id:    "777",
				title: "[Feature request] an awesome feature for `tool`",
			},
			want:    "777-feature-request-an-awesome-feature-for-tool",
			wantErr: false,
		},
		{
			name: "ending dot should be ignored",
			args: args{
				id:    "1234",
				title: "Bug when running something.",
			},
			want:    "1234-bug-when-running-something",
			wantErr: false,
		},
		{
			name: "branch name should be all lowercase",
			args: args{
				id:    "2345",
				title: "THIS IS AN IMPROVEMENT SUGGESTION :-) !!!",
			},
			want:    "2345-this-is-an-improvement-suggestion",
			wantErr: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateName(tt.id, tt.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("got err: %v, but wanted err set to %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
