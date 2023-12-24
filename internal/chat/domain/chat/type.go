package chat

type Type uint8

const (
	Direct Type = 1
	Group  Type = 2
)

func (t Type) Uint8() uint8 {
	return uint8(t)
}

func (t Type) Types() []Type {
	return []Type{Direct, Group}
}

func NewType(chatType uint8) (Type, error) {
	return Type(chatType), nil
}
