package greeting

import (
	"fmt"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		hour        string
		want        string
		expectedErr bool
	}{
		{"10:00", "こんにちは", false},
		// ...
	}
	for _, tt := range tests {
		t.Run(caseName(t, tt.hour, tt.want, tt.expectedErr), func(t *testing.T) {
			if got := Do(ToDate(t, tt.hour)); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func caseName(t *testing.T, time, want string, expectErr bool) string {
	t.Helper()
	if expectErr {
		return fmt.Sprintf("%sは時間外なのでerrと%sを返す", time, want)
	} else {
		t := ToDate(t, time).Hour()
		switch {
		case t >= 4 && t <= 9:
			return fmt.Sprintf("%sは朝なので%sを返す", time, want)
		case t >= 10 && t <= 16:
			return fmt.Sprintf("%sは昼なので%sを返す", time, want)
		default:
			return fmt.Sprintf("%sは夜なので%sを返す", time, want)
		}
	}
}

func ToDate(t *testing.T, date string) time.Time {
	t.Helper()
	d, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("2022-02-01 %s:00", date))
	if err != nil {
		t.Fatalf("toDate: %v", err)
	}
	return d
}
