package gomesh_test

import (
	"io"
	"os"
	"testing"

	"github.com/crhntr/gomesh"
)

func TestParseQualifierRecordSet(t *testing.T) {
	f, err := os.Open("testdata/qual2017.xml")
	if err != nil {
		t.Error(err)
	}

	qrc, errc := gomesh.ParseQualifierRecordSet(f)

	count := 0
	for {
		select {
		case qr := <-qrc:
			t.Logf("%d: %q", count, qr.UI)
			count++
		case err := <-errc:
			t.Logf("count: %d", count)
			if err == io.EOF {
				close(errc)
				return
			}
			t.Error(err)
			close(errc)
			return
		}
	}
}