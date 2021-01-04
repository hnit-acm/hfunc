package basic

import "hash"

type PasswordString String

func (p PasswordString) GetNative() string {
	return string(p)
}

type Password interface {
	Encode(hash hash.Hash, args ...string) string
}

func (p PasswordString) Encode(hash hash.Hash, args ...string) string {
	return p.Encode(hash, args...)
}
