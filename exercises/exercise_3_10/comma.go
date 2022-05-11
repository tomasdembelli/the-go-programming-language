package exercise310

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	sB := []byte(s)
	step := 3
	m := len(sB) % step

	var err error

	var buf2 bytes.Buffer
	for i := m; i < len(sB); i += step {
		_ = buf2.WriteByte(',')
		_, err = buf2.Write(sB[i : i+step])
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}

	if m != 0 {
		var buf bytes.Buffer
		_, err = buf.Write(append(sB[:m], buf2.Bytes()...))
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return buf.String()
	}

	return string(buf2.Bytes()[1:])
}
