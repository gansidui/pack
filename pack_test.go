package pack

import (
	"testing"
)

func TestPack(t *testing.T) {
	p := New()
	if p.Len() != 0 {
		t.Fatal()
	}

	p.WriteUint32(3200)
	p.WriteUint16(1600)
	p.WriteUint64(64)
	p.WriteUint8(8)
	p.Write([]byte("hello"))
	p.WriteString("world")

	if p.Len() != 25 {
		t.Fatal()
	}

	if v, ok := p.ReadUint32(); !ok || v != 3200 {
		t.Fatal()
	}
	if v, ok := p.ReadUint16(); !ok || v != 1600 {
		t.Fatal()
	}
	if v, ok := p.ReadUint64(); !ok || v != 64 {
		t.Fatal()
	}
	if v, ok := p.ReadUint8(); !ok || v != 8 {
		t.Fatal()
	}

	buf := make([]byte, 10)
	p.Read(buf[0:5])
	if string(buf[0:5]) != "hello" {
		t.Fatal()
	}

	p.Read(buf[5:10])
	if string(buf[5:10]) != "world" {
		t.Fatal()
	}

	if p.Len() != 0 {
		t.Fatal()
	}

	p.WriteUint16(1314)
	p.Reset()

	if p.Len() != 0 {
		t.Fatal()
	}

	p.Write([]byte("hello world"))
	if string(p.Bytes()) != "hello world" {
		t.Fatal()
	}

	buf = p.Next(5)
	if string(buf) != "hello" || p.Len() != 6 {
		t.Fatal()
	}

	buf = p.Next(6)
	if string(buf) != " world" || p.Len() != 0 {
		t.Fatal()
	}
}

func TestConvert(t *testing.T) {
	var a uint16 = 65535
	var b uint32 = 12345678
	var c uint64 = 12345678901234567890

	var buf []byte

	buf = Uint16ToBytes(a)
	if BytesToUint16(buf) != a {
		t.Fatal()
	}

	buf = Uint32ToBytes(b)
	if BytesToUint32(buf) != b {
		t.Fatal()
	}

	buf = Uint64ToBytes(c)
	if BytesToUint64(buf) != c {
		t.Fatal()
	}
}
