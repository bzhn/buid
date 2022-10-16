package file

import "testing"

func TestSizeMod(t *testing.T) {
	type test struct {
		input int64
		want  string
	}

	tt := []test{
		test{
			input: 4061287350129025369,
			want:  "9",
		},
		test{
			input: 2157025523395,
			want:  "3",
		},
		test{
			input: 32158,
			want:  "e",
		},
		test{
			input: 827,
			want:  "b",
		},
		test{
			input: 3,
			want:  "3",
		},
		test{
			input: 0,
			want:  "0",
		},
		test{
			input: 1705820,
			want:  "c",
		},
		test{
			input: 10000,
			want:  "0",
		},
	}

	for _, tc := range tt {
		if have := SizeMod(tc.input); have != tc.want {
			t.Logf("ERROR: want %s, got %s, for the %d file size", tc.want, have, tc.input)
			t.Fail()
		}
	}
}
func TestSizeApprox(t *testing.T) {
	type test struct {
		input int64
		want  string
	}

	tt := []test{
		test{
			input: 4061287350129025369,
			want:  "f",
		},
		test{
			input: 2157025523395,
			want:  "c",
		},
		test{
			input: 32158,
			want:  "4",
		},
		test{
			input: 827,
			want:  "2",
		},
		test{
			input: 3,
			want:  "0",
		},
		test{
			input: 0,
			want:  "0",
		},
		test{
			input: 1705820,
			want:  "6",
		},
		test{
			input: 10000,
			want:  "4",
		},
	}

	for _, tc := range tt {
		if have := SizeApprox(tc.input); have != tc.want {
			t.Logf("ERROR: want %s, got %s, for the %d file size", tc.want, have, tc.input)
			t.Fail()
		}
	}
}
