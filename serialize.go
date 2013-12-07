package driver

import (
	"encoding/binary"
	"github.com/tarm/goserial"
	"io"
)

type Connection struct {
	rwc io.ReadWriteCloser
}

func Connect() (c *Connection, err error) {
	c = &Connection{}
	c0 := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}

	c.rwc, err = serial.OpenPort(c0)
	if err != nil {
		return nil, err
	}

	return
}

func (c *Connection) TurnSingle(motorIndex byte, stepCount int16) {
	c.rwc.Write([]byte("t"))
	c.rwc.Write([]byte{motorIndex})
	binary.Write(c.rwc, binary.LittleEndian, stepCount)
}

func (c *Connection) TurnDouble(firstSteps int16, secondSteps int16) {
	c.rwc.Write([]byte("i"))
	binary.Write(c.rwc, binary.LittleEndian, firstMotorStepCount)
	binary.Write(c.rwc, binary.LittleEndian, secondMotorStepCount)
}
