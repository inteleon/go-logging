package helper

import (
	"fmt"
	"strings"
	"testing"
)

// TestBuffer
func TestBuffer(t *testing.T) {
	logging, w := NewTestLogging()

	for i := 0; i < 4; i++ {
		ii := i
		extra := ""
		for j := 0; j < i; j++ {
			extra = extra + "X"
		}
		logging.Info(fmt.Sprintf("MSG %d %s", ii, extra))
		// time.Sleep(time.Millisecond * 10)
	}

	if len(w.Buffer) != 4 {
		t.Fatalf("expected 4 entries")
	}
	if !strings.Contains(string(w.Buffer[0]), "MSG 0 ") {
		t.Errorf("expected msg 0 to contain MSG 0 was: %v", string(w.Buffer[0]))
	}
	if !strings.Contains(string(w.Buffer[1]), "MSG 1 X") {
		t.Errorf("expected msg 1 to contain MSG 1 X was: %v", string(w.Buffer[1]))
	}
	if !strings.Contains(string(w.Buffer[2]), "MSG 2 XX") {
		t.Errorf("expected msg 2 to contain MSG 2 XX was: %v", string(w.Buffer[2]))
	}
	if !strings.Contains(string(w.Buffer[3]), "MSG 3 XXX") {
		t.Errorf("expected msg 3 to contain MSG 3 XXX was: %v", string(w.Buffer[3]))
	}

}
