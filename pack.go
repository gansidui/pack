package pack

import (
	"bytes"
	"encoding/binary"
)

type Pack struct {
	buf *bytes.Buffer
}

func New() *Pack {
	return &Pack{
		buf: new(bytes.Buffer),
	}
}

func NewWithNoCopy(buf []byte) *Pack {
	return &Pack{
		buf: bytes.NewBuffer(buf),
	}
}

func (p *Pack) Reset() {
	p.buf.Reset()
}

func (p *Pack) Len() int {
	return p.buf.Len()
}

func (p *Pack) Bytes() []byte {
	return p.buf.Bytes()
}

func (p *Pack) Next(n int) []byte {
	return p.buf.Next(n)
}

func (p *Pack) Read(buf []byte) (n int, err error) {
	return p.buf.Read(buf)
}

func (p *Pack) ReadUint8() (uint8, bool) {
	buf := p.buf.Next(1)
	if len(buf) != 1 {
		return 0, false
	}
	return buf[0], true
}

func (p *Pack) ReadUint16() (uint16, bool) {
	buf := p.buf.Next(2)
	if len(buf) != 2 {
		return 0, false
	}
	return binary.BigEndian.Uint16(buf), true
}

func (p *Pack) ReadUint32() (uint32, bool) {
	buf := p.buf.Next(4)
	if len(buf) != 4 {
		return 0, false
	}
	return binary.BigEndian.Uint32(buf), true
}

func (p *Pack) ReadUint64() (uint64, bool) {
	buf := p.buf.Next(8)
	if len(buf) != 8 {
		return 0, false
	}
	return binary.BigEndian.Uint64(buf), true
}

func (p *Pack) Write(b []byte) (n int, err error) {
	return p.buf.Write(b)
}

func (p *Pack) WriteString(s string) (n int, err error) {
	return p.buf.WriteString(s)
}

func (p *Pack) WriteUint8(v uint8) {
	binary.Write(p.buf, binary.BigEndian, v)
}

func (p *Pack) WriteUint16(v uint16) {
	binary.Write(p.buf, binary.BigEndian, v)
}

func (p *Pack) WriteUint32(v uint32) {
	binary.Write(p.buf, binary.BigEndian, v)
}

func (p *Pack) WriteUint64(v uint64) {
	binary.Write(p.buf, binary.BigEndian, v)
}

func Uint64ToBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func Uint32ToBytes(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

func Uint16ToBytes(v uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
}

func BytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func BytesToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func BytesToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}
