package gext

import "bytes"

func writeID(cp component) string {
	var buffer bytes.Buffer

	if cp.GetID() != "" {
		buffer.WriteString(" id=")
		buffer.WriteString(`"`)
		buffer.WriteString(cp.GetID())
		buffer.WriteString(`" `)
	}
	return buffer.String()
}
