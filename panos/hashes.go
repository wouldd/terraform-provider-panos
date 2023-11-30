package panos

import (
	"bytes"
	"fmt"
	"hash/crc32"
)

func resourceMatchAddressPrefixHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s%t", m["prefix"].(string), m["exact"].(bool)))
	return HashString(buf.String())
}

func resourceTargetHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["serial"].(string))
	vl := m["vsys_list"].([]interface{})
	for i := range vl {
		buf.WriteString(vl[i].(string))
	}
	return HashString(buf.String())
}

func HashString(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
