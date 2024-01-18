package apple

import (
	"testing"

	"github.com/creepitall/goworm/internal/models"
)

func TestApple_IsCrossed(t *testing.T) {
	type fields struct {
		positions models.Positions
	}
	type args struct {
		p models.Position
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "value to first position ",
			fields: fields{
				positions: []models.Position{
					{X: 5, Y: 5},
					{X: 5, Y: 6},
				},
			},
			args: args{models.Position{X: 5, Y: 5}},
			want: fields{
				positions: []models.Position{
					{X: 5, Y: 6},
				},
			},
		},
		{
			name: "value to last position ",
			fields: fields{
				positions: []models.Position{
					{X: 5, Y: 5},
					{X: 5, Y: 6},
				},
			},
			args: args{models.Position{X: 5, Y: 6}},
			want: fields{
				positions: []models.Position{
					{X: 5, Y: 5},
				},
			},
		},
		{
			name: "value to mid position ",
			fields: fields{
				positions: []models.Position{
					{X: 5, Y: 5},
					{X: 5, Y: 6},
					{X: 5, Y: 7},
				},
			},
			args: args{models.Position{X: 5, Y: 6}},
			want: fields{
				positions: []models.Position{
					{X: 5, Y: 5},
					{X: 5, Y: 7},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Apple{
				positions: tt.fields.positions,
			}
			_ = a.IsCrossed(tt.args.p)
			if !toEqual(a.positions, tt.want.positions) {
				t.Errorf("Apple.IsCrossed() = %v, want %v", a.positions, tt.want.positions)
			}
		})
	}
}

func toEqual(one, two models.Positions) bool {
	if len(one) != len(two) {
		return false
	}

	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			return false
		}
	}

	return true
}
