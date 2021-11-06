package datareader

import (
	"io"
	"log"
)

// Reader implements io.Reader
var _ io.Reader = &Reader{}

// Reader represents bit-level data reader
type Reader struct {
	data        []byte
	length      int
	bitPosition uint64
	err         error
}

// NewReader creates new Reader
func NewReader(data []byte) *Reader {
	result := &Reader{
		data:   data,
		length: len(data),
	}

	return result
}

// Read implements io.Reader interface
func (r *Reader) Read(p []byte) (n int, err error) {
	if p == nil {
		return 0, nil
	}

	for ; n < len(p); n++ {
		result := r.GetByte()

		if err := r.Error(); err != nil {
			return n, err
		}

		p[n] = result
	}

	return n, nil
}

// SkipBits skips bits `count` bits
func (r *Reader) SkipBits(count uint64) {
	r.bitPosition += count
}

// SkipByte skips one byte (8 bits)
func (r *Reader) SkipByte() {
	r.SkipBits(byteLen)
}

// SkipBytes skips count bytes (count * 8 bits)
func (r *Reader) SkipBytes(count uint64) {
	r.SkipBits(count * byteLen)
}

// Align byte-aligns reader
func (r *Reader) Align() {
	offset := r.GetBitOffset()

	// no need to align anything
	if offset == 0 {
		return
	}

	r.SkipBits(uint64(byteLen - r.GetBitOffset()))
}

func (r *Reader) getCurrentByte() byte {
	pos := r.bitPosition / byteLen
	if pos >= uint64(r.length) {
		r.err = io.EOF
		return 0
	}

	result := r.data[pos]

	return result
}

// GetBitOffset returns current bit offset
func (r *Reader) GetBitOffset() byte {
	return byte(r.bitPosition % byteLen)
}

func (r *Reader) getBit() byte {
	result := ((r.getCurrentByte() >> r.GetBitOffset()) & bitMask)
	r.bitPosition++

	return result
}

// GetBit returns a single bit true/false
func (r *Reader) GetBit() bool {
	return r.getBit() > 0
}

func (r *Reader) getBits(count, limit int) (result uint64) {
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

// GetBits returns `count` bits as uint8
func (r *Reader) GetBits(count int) byte {
	return byte(r.getBits(count, byteLen))
}

// GetBits16 returns `count` bits as uint16
func (r *Reader) GetBits16(count int) uint16 {
	return uint16(r.getBits(count, int16Len))
}

// GetBits32 returns `count` bits as uint32
func (r *Reader) GetBits32(count int) uint32 {
	return uint32(r.getBits(count, int32Len))
}

// GetBits64 returns `count` bits as uint64
func (r *Reader) GetBits64(count int) uint64 {
	return r.getBits(count, int64Len)
}

// GetByte returns 8-bits byte
func (r *Reader) GetByte() byte {
	return r.GetBits(byteLen)
}

// GetBytes returns slice of `count` bytes
func (r *Reader) GetBytes(count int) (result []byte) {
	if count == 0 {
		return []byte{}
	}

	for i := 0; i < count; i++ {
		result = append(result, r.GetByte())
	}

	return result
}

// GetInt8 returns 8-bits signed int
func (r *Reader) GetInt8() int8 {
	return int8(r.GetByte())
}

// GetUint16 returns 16-bit unsigned int
func (r *Reader) GetUint16() uint16 {
	return r.GetBits16(int16Len)
}

// GetInt16 returns 16-bit signed int
func (r *Reader) GetInt16() int16 {
	return int16(r.GetUint16())
}

// GetUint32 returns 32-bit unsigned int
func (r *Reader) GetUint32() uint32 {
	return r.GetBits32(int32Len)
}

// GetInt32 returns 32-bit signed int
func (r *Reader) GetInt32() int32 {
	return int32(r.GetUint32())
}

// GetUint64 returns 64-bit unsigned int
func (r *Reader) GetUint64() uint64 {
	return r.GetBits64(int64Len)
}

// GetInt64 returns 64-bit signed int
func (r *Reader) GetInt64() int64 {
	return int64(r.GetUint64())
}

// Error returns not nil error if some errror reached during reading
// (the most likely it is io.EOF)
func (r *Reader) Error() error {
	return r.err
}
