package convert

func intToXlsxAxis(n int) string {
	buf := make([]byte, 0)
	q := int('Z' - 'A' + 1)
	for {
		buf = append([]byte{byte(n%q) + 'A'}, buf...)
		if n < q {
			break
		}
		n = (n - q) / q
	}

	return string(buf)
}
