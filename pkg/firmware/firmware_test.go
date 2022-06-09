package firmware_test

import (
	"testing"

	"github.com/pinheirolucas/nps/pkg/firmware"
	"github.com/stretchr/testify/assert"
)

func TestFirmware_String(t *testing.T) {
	testCases := []struct {
		name  string
		fwStr string
		want  string
	}{
		{
			name:  "left pad on minors",
			fwStr: "1.01",
			want:  "1.01",
		},
		{
			name:  "no left pad",
			fwStr: "1.11",
			want:  "1.11",
		},
		{
			name:  "v0",
			fwStr: "0.00",
			want:  "0.00",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(st *testing.T) {
			fw, err := firmware.Parse(testCase.fwStr)
			if err != nil {
				st.Fatal(err)
			}

			got := fw.String()
			assert.Equal(st, got, testCase.want)
		})
	}
}
