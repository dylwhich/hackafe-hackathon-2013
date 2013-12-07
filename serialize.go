package driver

import (
	"encoding/binary"
	"github.com/tarm/goserial"
	"io"
)

const (
	Baud = 9600
)

type Connection struct {
	rwc io.ReadWriteCloser
}

func Connect(device string) (c *Connection, err error) {
	c = &Connection{}
	c0 := &serial.Config{
		Name: device,
		Baud: Baud,
	}

	c.rwc, err = serial.OpenPort(c0)
	if err != nil {
		return nil, err
	}

	return
}

func (c *Connection) StepSingle(motorIndex byte, stepCount int16) {
	c.rwc.Write([]byte("t"))
	c.rwc.Write([]byte{motorIndex})
	binary.Write(c.rwc, binary.LittleEndian, stepCount)
}

func (c *Connection) StepDouble(firstSteps int16, secondSteps int16) {
	c.rwc.Write([]byte("i"))
	binary.Write(c.rwc, binary.LittleEndian, firstSteps)
	binary.Write(c.rwc, binary.LittleEndian, secondSteps)
}
