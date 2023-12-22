package hash

import (
	"hash"
	"hash/crc32"
	"io"
)

func Hash(r io.Reader) []byte {
	hashier := dst_hash{
		hashier: crc32.NewIEEE(),
	}
	io.Copy(hashier, r)

	return hashier.hashier.Sum(nil)
}

type dst_hash struct {
	hashier hash.Hash32
}

func (dst dst_hash) Write(b []byte) (int, error) {
	return dst.hashier.Write(b)
}
