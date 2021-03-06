package diadem

import (
	"github.com/diademnetwork/go-diadem/util"
)

func PermPrefix(addr Address) []byte {
	return util.PrefixKey(contractPrefix(addr), []byte("permission"))
}

func contractPrefix(addr Address) []byte {
	return util.PrefixKey([]byte("contract"), []byte(addr.Local))
}

func DataPrefix(addr Address) []byte {
	return util.PrefixKey(contractPrefix(addr), []byte("data"))
}

func TextKey(addr Address) []byte {
	return util.PrefixKey(contractPrefix(addr), []byte("text"))
}
