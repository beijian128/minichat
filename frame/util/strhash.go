package util

var crcTable = make([]uint32, 256)

const crcPOLY uint32 = 0x04c11db7

var crcTableInitialized bool

func init() {
	initCRCTable()
}

func initCRCTable() {
	if crcTableInitialized {
		return
	}
	var (
		i uint32
		c uint32
		j uint32
	)
	for i = 0; i < 256; i++ {
		c = i << 24
		for j = 8; j != 0; j = j - 1 {
			if (c & 0x80000000) != 0 {
				c = (c << 1) ^ crcPOLY
			} else {
				c = c << 1
			}
			crcTable[i] = c
		}
	}
	crcTableInitialized = true
}

// StringHash 字符串转为32位整形值
func StringHash(s string) uint32 {
	var (
		hash uint32
		b    uint32
	)
	for _, c := range s {
		b = uint32(c)
		hash = ((hash >> 8) & 0x00FFFFFF) ^ crcTable[(hash^b)&0x000000FF]
	}
	return hash
}
