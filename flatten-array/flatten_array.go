package flatten

func Flatten(nested interface{}) []interface{} {
	res := make([]interface{}, 0)
	for _, v := range nested.([]interface{}) {
		if v == nil {
			continue
		}
		switch v.(type) {
		case []interface{}:
			res = append(res, Flatten(v)...)
		default:
			res = append(res, v)
		}
	}
	return res
}
