package zimbraCommon

type ZBool bool

func (zb ZBool) MarshalJSON() ([]byte, error) {
	if zb {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

func (zb *ZBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "1" || s == `"1"` {
		*zb = true
	} else {
		*zb = false
	}
	return nil
}
