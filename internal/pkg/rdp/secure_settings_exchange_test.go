package rdp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClientInfoPDU_Serialize(t *testing.T) {
	req := NewClientInfoPDU("192.168.1.2", "User", "p@$$w0rd")

	expected := []byte{
		0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x53, 0x01, 0x01, 0x00, 0x16, 0x00, 0x08, 0x00,
		0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x31, 0x00, 0x39, 0x00, 0x32, 0x00, 0x2e, 0x00, 0x31, 0x00,
		0x36, 0x00, 0x38, 0x00, 0x2e, 0x00, 0x31, 0x00, 0x2e, 0x00, 0x32, 0x00, 0x00, 0x00, 0x55, 0x00,
		0x73, 0x00, 0x65, 0x00, 0x72, 0x00, 0x00, 0x00, 0x70, 0x00, 0x40, 0x00, 0x24, 0x00, 0x24, 0x00,
		0x77, 0x00, 0x30, 0x00, 0x72, 0x00, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	actual := req.Serialize()

	require.Equal(t, expected, actual)
}
