package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getOrderByStatus(t *testing.T) {
	type args struct {
		order []Order
		want  string
	}
	tests := []struct {
		name string
		args args
		want []Order
	}{
		{
			name: "good response",
			args: args{order: []Order{
				{
					Id:           "abcdef-123426",
					Status:       "Pending_Invoice",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123428",
					Status:       "shipped",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123427",
					Status:       "Pending_Invoice",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123429",
					Status:       "completed",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
			},
				want: "Pending_Invoice",
			},
			want: []Order{
				{
					Id:           "abcdef-123426",
					Status:       "Pending_Invoice",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123427",
					Status:       "Pending_Invoice",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123428",
					Status:       "shipped",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123429",
					Status:       "completed",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getOrderByStatus(tt.args.order, tt.args.want)
			assert.Equal(t, tt.want, got, "not getting expected response")
		})
	}
}

func Test_getOrderById(t *testing.T) {
	type args struct {
		order []Order
		want  string
	}
	tests := []struct {
		name string
		args args
		want Order
	}{
		{
			name: "good response",
			args: args{order: []Order{
				{
					Id:           "abcdef-123126",
					Status:       "Pending_Invoice",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123496",
					Status:       "shipped",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123426",
					Status:       "Pending_Invoice",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
				{
					Id:           "abcdef-123427",
					Status:       "completed",
					Items:        []Items{},
					Total:        12,
					CurrencyUnit: "usd"},
			},
				want: "abcdef-123426",
			},
			want: Order{

				Id:           "abcdef-123426",
				Status:       "Pending_Invoice",
				Items:        []Items{},
				Total:        12,
				CurrencyUnit: "usd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getOrderById(tt.args.order, tt.args.want)
			assert.Equal(t, tt.want, got, "not getting expected response")
		})
	}
}
