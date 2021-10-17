package datareader

import (
	"log"
)

type Reader struct {
	data        []byte
	bitPosition uint64
}

func NewReader(data []byte) *Reader {
	result := &Reader{
		data: data,
	}

	return result
}

func (r *Reader) SkipBits(count uint64) {
	r.bitPosition += count
}

func (r *Reader) SkipByte() {
	r.SkipBits(byteLen)
}

func (r *Reader) SkipBytes(count uint64) {
	r.SkipBits(count * byteLen)
}

func (r *Reader) Align() {
	r.SkipBits(uint64(byteLen - r.GetBitOffset()))
}

func (r *Reader) getCurrentByte() byte {
	return r.data[r.bitPosition/byteLen]
}

func (r *Reader) GetBitOffset() byte {
	return byte(r.bitPosition % byteLen)
}

func (r *Reader) getBit() byte {
	result := ((r.getCurrentByte() >> r.GetBitOffset()) & bitMask)
	r.bitPosition++

	return result
}

func (r *Reader) GetBit() bool {
	return r.getBit() > 0
}

func (r *Reader) getBits(count int, limit int) (result uint64) {
	switch {
	case count == 0:
		return 0
	case count > limit:
		log.Panicf("datareader.Reader: getBits(%d > %d)", count, limit)
	}

	for i := 0; i < count; i++ {
		result |= uint64(r.getBit()) << i
	}

	return result
}

func (r *Reader) GetBits(count int) byte {
	return byte(r.getBits(count, byteLen))
}

func (r *Reader) GetBits16(count int) uint16 {
	return uint16(r.getBits(count, int16Len))
}

func (r *Reader) GetBits32(count int) uint32 {
	return uint32(r.getBits(count, int32Len))
}

func (r *Reader) GetBits64(count int) uint64 {
	return r.getBits(count, int64Len)
}

func (r *Reader) GetByte() byte {
	return r.GetBits(byteLen)
}

func (r *Reader) GetBytes(count int) (result []byte) {
	for i := 0; i < count; i++ {
		result = append(result, r.GetByte())
	}

	return result
}

func (r *Reader) GetInt8() int8 {
	return int8(r.GetByte())
}

func (r *Reader) GetUint16() uint16 {
	return r.GetBits16(int16Len)
}

func (r *Reader) GetInt16() int16 {
	return int16(r.GetUint16())
}

func (r *Reader) GetUint32() uint32 {
	return r.GetBits32(int32Len)
}

func (r *Reader) GetInt32() int32 {
	return int32(r.GetUint32())
}
