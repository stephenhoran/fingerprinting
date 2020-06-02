package fingerprinting

import (
	"hash/fnv"
)

const SepByte byte = 255

func NewGoHash(labels map[string]string) uint64 {
	h := fnv.New64a()

	if len(labels) == 0 {
		return h.Sum64()
	}

	for labelName, labelValue := range labels {
		h.Write([]byte(labelName))
		h.Write([]byte{SepByte})
		h.Write([]byte(labelValue))
		h.Write([]byte{SepByte})
	}

	return h.Sum64()
}
