package bestring

import (
	"hash/crc32"
)

func checkSum(strBytes []byte) uint32 {
	return crc32.ChecksumIEEE(strBytes)
}

/* func apHash(strBytes []byte) (hash uint32) {
	for i := 0; i < len(strBytes); i++ {
		if (i & 1) == 0 {
			hash ^= ((hash << 7) ^ uint32(strBytes[i]) ^ (hash >> 3))
		} else {
			hash ^= (^((hash << 11) ^ uint32(strBytes[i]) ^ (hash >> 5)) + 1)
		}
	}

	return hash & 0x7FFFFFFF
}

func md5Sum(any interface{}) string {
	switch val := any.(type) {
	case []byte:
		return fmt.Sprintf("%x", md5.Sum(val))
	case string:
		return fmt.Sprintf("%x", md5.Sum([]byte(val)))
	default:
		h := md5.New()
		fmt.Fprintf(h, "%v", val)
		return fmt.Sprintf("%x", h.Sum(nil))
	}
}
*/
