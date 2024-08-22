package patchway

type Operation interface {
	String() string
	Bytes() []byte
}
