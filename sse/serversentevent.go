package sse

import (
	"bytes"
	"fmt"
	"io"
)

type Event struct {
	Data []byte
}

func (ev *Event) MarshalTo(w io.Writer) error {
	if len(ev.Data) == 0 {
		return nil
	}

	if len(ev.Data) > 0 {
		sd := bytes.Split(ev.Data, []byte("\n"))
		for i := range sd {
			if _, err := fmt.Fprintf(w, "data: %s\n", sd[i]); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprint(w, "\n"); err != nil {
		return err
	}

	return nil
}
