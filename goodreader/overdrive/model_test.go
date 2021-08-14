package goodreader

import (
	"reflect"
	"testing"
)

func Test_extract(t *testing.T) {
	type args struct {
		jsonStart []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *[]BookSearchResultJson
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				jsonStart: []byte(`[
				{
					"firstCreatorName": "the author",
					"sortTitle": "the title",
					"type": {"id": "ebook"},
					"isAvailable": true
				}
				]`),
			},
			want: &[]BookSearchResultJson{
				{
					Author:      "the author",
					Title:       "the title",
					IsAvailable: true,
					ItemType: ItemType{
						Id: "ebook",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "malformed - missing closing square bracket",
			args: args{
				jsonStart: []byte(`[
				{
					"firstCreatorName": "the author",
					"sortTitle": "the title",
					"type": {"id": "ebook"},
					"isAvailable": true
				}	
				`),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extract(tt.args.jsonStart)
			if (err != nil) != tt.wantErr {
				t.Errorf("extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extract() = %v, want %v", got, tt.want)
			}
		})
	}
}
