package enumsx

type EnumString string

func (e EnumString) Id() string {
	return string(e)
}

type EnumInt int

func (e EnumInt) Id() int {
	return int(e)
}

type EnumInt8 int8

func (e EnumInt8) Id() int8 {
	return int8(e)
}

type EnumInt16 int16

func (e EnumInt16) Id() int16 {
	return int16(e)
}

type EnumInt32 int32

func (e EnumInt32) Id() int32 {
	return int32(e)
}

type EnumUint uint

func (e EnumUint) Id() uint {
	return uint(e)
}

type EnumUint8 uint8

func (e EnumUint8) Id() uint8 {
	return uint8(e)
}

type EnumUint16 uint16

func (e EnumUint16) Id() uint16 {
	return uint16(e)
}

type EnumUint32 uint32

func (e EnumUint32) Id() uint32 {
	return uint32(e)
}
