package serializer

import (
	"github.com/tarm/goserial"
	"io"
	"unsafe"
)


type Connection struct {
	inner io.ReadWriteCloser
}

func Connect() (*Connection, error) {
	c0 := &serial.Config{Name:"/dev/ttyUSB0", Baud: 9600}

	in, e := serial.OpenPort(c0)

	if e != nil {
		return nil, e
	}

	connection := &Connection{inner: in}

	return connection, nil
}

func (c *Connection) TurnSingle(motorIndex byte, stepCount int16) {
	c.inner.Write([]byte('t'))
	c.inner.Write([]byte(motorIndex))
	c.inner.Write([]byte(stepCount))
}

func (c *Connection) TurnDouble(firstMotorStepCount int16, secondMotorStepCount int16){
	c.inner.Write([]byte('i'))
	c.inner.Write([]byte(firstMotorStepCount))
	c.inner.Write([]byte(secondMotorStepCount))
}
