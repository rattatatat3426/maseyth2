package wire

import (
	"encoding/binary"
	"testing"

	"github.com/rattatatat3426/maseyth2/internal/protocol"
	"github.com/rattatatat3426/maseyth2/quicvarint"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWire(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wire Suite")
}

func encodeVarInt(i uint64) []byte {
	return quicvarint.Append(nil, i)
}

func appendVersion(data []byte, v protocol.VersionNumber) []byte {
	offset := len(data)
	data = append(data, []byte{0, 0, 0, 0}...)
	binary.BigEndian.PutUint32(data[offset:], uint32(v))
	return data
}
