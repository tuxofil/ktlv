package ktlv

import (
	"io"
)

type KV struct {
	Key   Key
	Value interface{}
}

func Write(writer io.Writer, types []FType, elements ...KV) (int, error) {
	var (
		written int
		elem    *Elem
	)
	for _, kv := range elements {
		elem = &Elem{
			Key:   kv.Key,
			FType: types[int(kv.Key)],
			Value: kv.Value,
		}
		n, err := elem.WriteTo(writer)
		written += n
		if err != nil {
			return written, err
		}
	}
	return written, nil
}
