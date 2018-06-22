package util

//过滤字段，map中只保留需要的键值对
func OnlyCols(cols []string, data *map[string]string) {
	for k := range *data {
		//判断k 是否在需要的cols中， 如果不在，则对应的键值对
		have := false //不在
		for _, col := range cols {
			if k == col {
				have = true
				break
			}
		}
		if have == false {
			delete(*data, k)
		}
	}
}
