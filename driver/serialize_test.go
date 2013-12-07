package driver

import (
	//	"bufio"
	"testing"
)

/*
type bufReadWriteCloser struct {
	closed bool
	*bufio.Reader
	*bufio.Writer
}

func (b *bufReadWriteCloser) Read(b []byte) (int, err) {
	if b.closed {
		return 0, nil
	}

	return b.Reader.Read(b)
}

func (b *bufReadWriteCloser) Write(b []byte) (int, err) {
	if b.closed {
		return 0, nil
	}

	return b.Writer.Write(b)
}

func (b *bufReadWriteCloser) Close() {
	b.closed = true
}*/

func TestStepSingleReal(t *testing.T) {
	return

	c, err := Connect("/dev/ttyUSB0")
	if err != nil {
		t.Fatalf("Could not connect to serial: %s\n", err)
	}

	t.Logf("Testing single step 400 times, motor 0\n")
	c.MoveRelativeSingle(0, 400)

	t.Logf("Testing single step 400 times, motor 1\n")
	c.MoveRelativeSingle(1, 400)

	t.Logf("Testing single step -400 times, motor 0\n")
	c.MoveRelativeSingle(0, -400)

	t.Logf("Testing single step -400 times, motor 1\n")
	c.MoveRelativeSingle(1, -400)
}

/*
func TestStepSingleLeft(t *testing.T) {
	c, err := Connect("/dev/ttyUSB0")
	if err != nil {
		t.Fatalf("Could not connect to serial: %s\n", err)
	}

	c.StepSingle(byte(1), int16(-2400))
}

func TestDoubleStepUp(t *testing.T) {
	return
	c, err := Connect("/dev/ttyUSB0")
	if err != nil {
		t.Fatalf("Could not connect to serial: %s\n", err)
	}

	c.StepDouble(int16(-800), int16(-800))
}
*/
