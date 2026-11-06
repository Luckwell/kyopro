package lib

import (
	"bufio"
	"io"
)

// Scanner provides zero-allocation reads directly from the byte stream.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner creates a new FastScanner with a 1MB buffer.
// Typical usage: in := fastio.NewScanner(os.Stdin)
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReaderSize(r, 1<<20)}
}

func (s *Scanner) readByte() byte {
	b, err := s.r.ReadByte()
	if err != nil {
		return 0
	}
	return b
}

func (s *Scanner) skipWhitespace() byte {
	b := s.readByte()
	for b == ' ' || b == '\n' || b == '\r' || b == '\t' {
		b = s.readByte()
	}
	return b
}

// Int reads the next space-delimited token as an int.
func (s *Scanner) Int() int {
	return int(s.Int64())
}

// Int64 reads the next space-delimited token as an int64.
func (s *Scanner) Int64() int64 {
	b := s.skipWhitespace()
	if b == 0 {
		return 0
	}

	sign := int64(1)
	if b == '-' {
		sign = -1
		b = s.readByte()
	}

	var val int64
	for b >= '0' && b <= '9' {
		val = val*10 + int64(b-'0')
		b = s.readByte()
	}
	return val * sign
}

// Float64 reads the next space-delimited token as a float64.
func (s *Scanner) Float64() float64 {
	b := s.skipWhitespace()
	if b == 0 {
		return 0
	}

	sign := 1.0
	if b == '-' {
		sign = -1.0
		b = s.readByte()
	}

	var val float64
	for b >= '0' && b <= '9' {
		val = val*10 + float64(b-'0')
		b = s.readByte()
	}

	if b == '.' {
		b = s.readByte()
		frac := 1.0
		for b >= '0' && b <= '9' {
			frac /= 10
			val += float64(b-'0') * frac
			b = s.readByte()
		}
	}
	return val * sign
}

// String reads the next space-delimited token as a string.
func (s *Scanner) String() string {
	return string(s.Bytes())
}

// Bytes reads the next space-delimited token as a byte slice.
func (s *Scanner) Bytes() []byte {
	b := s.skipWhitespace()
	if b == 0 {
		return nil
	}

	var res []byte
	for b != ' ' && b != '\n' && b != '\r' && b != '\t' && b != 0 {
		res = append(res, b)
		b = s.readByte()
	}
	return res
}

// Writer provides high-speed, buffered integer output.
type Writer struct {
	w *bufio.Writer
}

// NewWriter creates a new FastWriter with a 1MB buffer.
// Typical usage: out := fastio.NewWriter(os.Stdout); defer out.Flush()
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: bufio.NewWriterSize(w, 1<<20)}
}

// Flush writes any buffered data to the underlying io.Writer.
func (w *Writer) Flush() {
	w.w.Flush()
}

func (w *Writer) PrintByte(b byte) {
	w.w.WriteByte(b)
}

func (w *Writer) PrintString(s string) {
	w.w.WriteString(s)
}

// PrintInt writes an integer directly to the buffer.
func (w *Writer) PrintInt(n int) {
	w.PrintInt64(int64(n))
}

// PrintInt64 writes an int64 directly to the buffer.
func (w *Writer) PrintInt64(n int64) {
	if n == 0 {
		w.w.WriteByte('0')
		return
	}
	if n < 0 {
		w.w.WriteByte('-')
		n = -n
	}

	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	w.w.Write(buf[i:])
}
