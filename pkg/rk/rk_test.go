package rk

import (
	"reflect"
	"testing"
)

func TestRK_Permutation(t *testing.T) {
	tests := []struct {
		name string
		r    RK
		want []int
	}{
		{
			"test1",
			RK{1, 0.5, 0},
			[]int{2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Permutation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RK.Permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRK_Clone(t *testing.T) {
	tests := []struct {
		name string
		r    RK
		want RK
	}{
		{
			"test1",
			RK{1, 0.5, 0},
			RK{1, 0.5, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RK.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRK_Sort(t *testing.T) {
	tests := []struct {
		name string
		r    RK
		want RK
	}{
		{
			"test1",
			RK{1, 0.5, 0},
			RK{0, 0.5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RK.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMean(t *testing.T) {
	type args struct {
		rks []RK
	}
	tests := []struct {
		name string
		args args
		want RK
	}{
		{
			"test1",
			args{
				[]RK{
					RK{1, 0.5, 0},
					RK{1, 0.5, 0},
					RK{1, 0.5, 0},
				},
			},
			RK{1, 0.5, 0},
		},
		{
			"test2",
			args{
				[]RK{
					RK{0, 0.5, 1},
					RK{1, 0.5, 0},
				},
			},
			RK{0.5, 0.5, 0.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.rks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRK_Uniformize(t *testing.T) {
	tests := []struct {
		name string
		r    RK
		want RK
	}{
		{
			"test1",
			RK{-5.25, 0.1, 75.99},
			RK{0, 0.5, 1},
		},
		{
			"test2", // NOT SURE HOW YET
			RK{0, 0, 0},
			RK{0, 0.5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Uniformize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RK.Uniformize() = %v, want %v", got, tt.want)
			}
		})
	}
}
