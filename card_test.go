package main

//
//import (
//	"reflect"
//	"testing"
//	"time"
//)
//
//package card
//
//import (
//"reflect"
//"testing"
//"time"
//
//)
//
//func TestCard_IsDeleted(t *testing.T) {
//	type fields struct {
//		card Card
//	}
//
//	tests := map[string]struct {
//		fields fields
//		want   bool
//	}{
//		"it returns false if the DeletedAt is a zero Date": {
//			fields: fields{card: Card{}},
//			want:   false,
//		},
//		"it returns true if the DeletedAt is not a zero Date": {
//			fields: fields{card: Card{DeletedAt: time.Now()}},
//			want:   true,
//		},
//	}
//
//	for name, tt := range tests {
//		tt := tt
//
//		t.Run(name, func(t *testing.T) {
//			if got := tt.fields.card.IsDeleted(); got != tt.want {
//				t.Errorf("IsDeleted() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestCard_IsNew(t *testing.T) {
//	t.Parallel()
//
//	type fields struct {
//		card Card
//	}
//
//	tests := map[string]struct {
//		fields fields
//		want   bool
//	}{
//		"cards with id = 0 are considered new": {
//			fields: fields{card: Card{ID: 0}},
//			want:   true,
//		},
//		"cards with non-zero ID should not be considered new": {
//			fields: fields{card: Card{ID: 1}},
//			want:   false,
//		},
//	}
//
//	for name, tt := range tests {
//		tt := tt
//
//		t.Run(name, func(t *testing.T) {
//			t.Parallel()
//
//			if got := tt.fields.card.IsNew(); got != tt.want {
//				t.Errorf("IsDeleted() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestCard_Expired(t *testing.T) {
//	t.Parallel()
//
//	oneYear := 365 * 24 * time.Hour
//
//	tests := map[string]struct {
//		args time.Time
//		want bool
//	}{
//		"has not expired":    {args: time.Now().Add(oneYear), want: false},
//		"has expired":        {args: time.Now().Add(-oneYear), want: true},
//		"expires in 5 years": {args: time.Now().Add(5 * oneYear), want: false},
//		"expires today":      {args: time.Now().Add(time.Second), want: false},
//	}
//
//	for name, tt := range tests {
//		tt := tt
//
//		t.Run(name, func(t *testing.T) {
//			t.Parallel()
//
//			c := &Card{ExpiryDate: tt.args}
//
//			if got := c.IsExpired(); got != tt.want {
//				t.Errorf("IsExpired() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestCard_Ok(t *testing.T) {
//	type fields struct {
//		card Card
//	}
//
//	type want struct {
//		Err error
//	}
//
//	tests := map[string]struct {
//		fields fields
//		want   want
//	}{
//		"wont save a card that doesn't belong to the user": {
//			fields: fields{
//				card: Card{
//					UserID:     "farouq",
//					Token:      "some-token",
//					CardNumber: "7777777777777777",
//				},
//			},
//			want: want{Err: ErrCardEmptyCardName},
//		},
//		"wont save a card with card number less than 16 digits": {
//			fields: fields{
//				card: Card{
//					UserID:     "samir",
//					Token:      "some-token",
//					CardNumber: "12345",
//				},
//			},
//			want: want{Err: ErrCardInvalidCardNumber},
//		},
//		"wont save a card with no CardName": {
//			fields: fields{
//				card: Card{
//					UserID:     "samir",
//					Token:      "some-token",
//					CardNumber: "1234512345123456",
//				},
//			},
//			want: want{Err: ErrCardEmptyCardName},
//		},
//		"wont save a card with no token": {
//			fields: fields{
//				card: Card{
//					UserID:     "samir",
//					CardNumber: "1234123412341234",
//				},
//			},
//			want: want{Err: ErrCardEmptyToken},
//		},
//	}
//
//	for name, tt := range tests {
//		tt := tt
//
//		t.Run(name, func(t *testing.T) {
//			err := tt.fields.card.Ok()
//			assert.ErrorIs(t, err, tt.want.Err)
//
//			// Make sure all errors are validation errors.
//			assert.ErrorIs(t, err, errors.E().WithKind(errors.KindValidation))
//		})
//	}
//}
//
//func Test_extractExpiryDate(t *testing.T) {
//	type args struct {
//		exp string
//	}
//
//	tests := map[string]struct {
//		args args
//		want time.Time
//	}{
//		"it can convert a valid expiry": {
//			args: args{exp: "2512"},
//			want: time.Date(2025, time.Month(12)+1, 0, 23, 59, 59, 0, time.UTC),
//		},
//		"the base year is 2000": {
//			args: args{exp: "0012"},
//			want: time.Date(2000, time.Month(12)+1, 0, 23, 59, 59, 0, time.UTC),
//		},
//		"invalid month adds to the years": {
//			args: args{exp: "9913"},
//			want: time.Date(2100, time.Month(1)+1, 0, 23, 59, 59, 0, time.UTC),
//		},
//	}
//
//	for name, tt := range tests {
//		tt := tt
//
//		t.Run(name, func(t *testing.T) {
//			if got := extractExpiryDate(tt.args.exp); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("extractExpiryDate() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
