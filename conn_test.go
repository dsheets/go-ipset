package ipset

import (
	"net"
	"testing"

	"github.com/mdlayher/netlink"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ti-mo/netfilter"
)

type queryMock struct {
	mock.Mock
}

func (q queryMock) Query(nlm netlink.Message) ([]netlink.Message, error) {
	args := q.Called(nlm.Data)
	return args.Get(0).([]netlink.Message), args.Error(1)
}

func TestConn_Protocol(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00}
	m.On("Query", data).Return([]netlink.Message{
		{Data: []byte{0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00}},
	}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}

	res, err := c.Protocol()
	if assert2.NoError(err) {
		assert2.Equal(uint8(6), res.Protocol.Get())
	}

	m.AssertExpectations(t)
}

func TestConn_Create(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x66, 0x6f, 0x6f, 0x00, 0x0d, 0x00, 0x03, 0x00, 0x68, 0x61, 0x73, 0x68, 0x3a, 0x6d, 0x61, 0x63,
		0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x05, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x07, 0x80,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.Create("foo", "hash:mac", 0, 0))

	m.AssertExpectations(t)
}

func TestConn_DestroyOne(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x66, 0x6f, 0x6f, 0x00,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.Destroy("foo"))

	m.AssertExpectations(t)
}

func TestConn_DestroyAll(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.DestroyAll())

	m.AssertExpectations(t)
}

func TestConn_FlushOne(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x66, 0x6f, 0x6f, 0x00,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.Flush("foo"))

	m.AssertExpectations(t)
}

func TestConn_FlushAll(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.FlushAll())

	m.AssertExpectations(t)
}

func TestConn_Rename(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x62, 0x61, 0x72, 0x00, 0x08, 0x00, 0x03, 0x00, 0x62, 0x61, 0x7a, 0x00,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.Rename("bar", "baz"))

	m.AssertExpectations(t)
}

func TestConn_Swap(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x62, 0x61, 0x72, 0x00, 0x08, 0x00, 0x03, 0x00, 0x62, 0x61, 0x7a, 0x00,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.Swap("bar", "baz"))

	m.AssertExpectations(t)
}

func TestConn_Test(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x62, 0x61, 0x7a, 0x00, 0x10, 0x00, 0x07, 0x80, 0x0c, 0x00, 0x01, 0x80, 0x08, 0x00, 0x01, 0x40,
		0xc0, 0xa8, 0x01, 0x01,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.Test("baz", EntryIP(net.ParseIP("192.168.1.1"))))

	m.AssertExpectations(t)
}

func TestConn_Header(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x62, 0x61, 0x7a, 0x00,
	}
	m.On("Query", data).Return([]netlink.Message{
		{Data: []byte{
			0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
			0x62, 0x61, 0x7a, 0x00, 0x0c, 0x00, 0x03, 0x00, 0x68, 0x61, 0x73, 0x68, 0x3a, 0x69, 0x70, 0x00,
			0x05, 0x00, 0x05, 0x00, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00,
		}},
	}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}

	res, err := c.Header("baz")
	if assert2.NoError(err) {
		assert2.Equal("hash:ip", res.TypeName.Get())
	}

	m.AssertExpectations(t)
}

func TestConn_List(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00,
	}
	m.On("Query", data).Return([]netlink.Message{
		{Data: []byte{
			0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
			0x62, 0x61, 0x61, 0x00, 0x0c, 0x00, 0x03, 0x00, 0x68, 0x61, 0x73, 0x68, 0x3a, 0x69, 0x70, 0x00,
			0x05, 0x00, 0x05, 0x00, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x04, 0x00, 0x04, 0x00, 0x00, 0x00,
			0x2c, 0x00, 0x07, 0x80, 0x08, 0x00, 0x12, 0x40, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0x13, 0x40,
			0x00, 0x01, 0x00, 0x00, 0x08, 0x00, 0x19, 0x40, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x1a, 0x40,
			0x00, 0x00, 0x00, 0x58, 0x08, 0x00, 0x18, 0x40, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x08, 0x80,
		}},
		{Data: []byte{
			0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
			0x62, 0x61, 0x72, 0x00, 0x0d, 0x00, 0x03, 0x00, 0x68, 0x61, 0x73, 0x68, 0x3a, 0x6d, 0x61, 0x63,
			0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x04, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x07, 0x80, 0x08, 0x00, 0x12, 0x40, 0x00, 0x00, 0x04, 0x00,
			0x08, 0x00, 0x13, 0x40, 0x00, 0x01, 0x00, 0x00, 0x08, 0x00, 0x19, 0x40, 0x00, 0x00, 0x00, 0x00,
			0x08, 0x00, 0x1a, 0x40, 0x00, 0x00, 0x01, 0x18, 0x08, 0x00, 0x18, 0x40, 0x00, 0x00, 0x00, 0x03,
			0x34, 0x00, 0x08, 0x80, 0x10, 0x00, 0x07, 0x80, 0x0a, 0x00, 0x11, 0x00, 0x01, 0x23, 0x45, 0x67,
			0x89, 0xaf, 0x00, 0x00, 0x10, 0x00, 0x07, 0x80, 0x0a, 0x00, 0x11, 0x00, 0x01, 0x23, 0x45, 0x67,
			0x89, 0xae, 0x00, 0x00, 0x10, 0x00, 0x07, 0x80, 0x0a, 0x00, 0x11, 0x00, 0x01, 0x23, 0x45, 0x67,
			0x89, 0xad, 0x00, 0x00,
		}},
		{Data: []byte{
			0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
			0x62, 0x61, 0x7a, 0x00, 0x0c, 0x00, 0x03, 0x00, 0x68, 0x61, 0x73, 0x68, 0x3a, 0x69, 0x70, 0x00,
			0x05, 0x00, 0x05, 0x00, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x2c, 0x00, 0x07, 0x80, 0x08, 0x00, 0x12, 0x40, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0x13, 0x40,
			0x00, 0x01, 0x00, 0x00, 0x08, 0x00, 0x19, 0x40, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x1a, 0x40,
			0x00, 0x00, 0x00, 0xe8, 0x08, 0x00, 0x18, 0x40, 0x00, 0x00, 0x00, 0x03, 0x34, 0x00, 0x08, 0x80,
			0x10, 0x00, 0x07, 0x80, 0x0c, 0x00, 0x01, 0x80, 0x08, 0x00, 0x01, 0x00, 0xc0, 0xa8, 0x08, 0x03,
			0x10, 0x00, 0x07, 0x80, 0x0c, 0x00, 0x01, 0x80, 0x08, 0x00, 0x01, 0x00, 0xc0, 0xa8, 0x08, 0x02,
			0x10, 0x00, 0x07, 0x80, 0x0c, 0x00, 0x01, 0x80, 0x08, 0x00, 0x01, 0x00, 0xc0, 0xa8, 0x08, 0x01,
		}},
	}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}

	res, err := c.ListAll()
	if assert2.NoError(err) {
		assert2.Len(res, 3)

		p := &res[0]
		assert2.Equal("baa", p.Name.Get())
		assert2.Equal("hash:ip", p.TypeName.Get())
		assert2.Len(p.Entries, 0)

		p = &res[1]
		assert2.Equal("bar", p.Name.Get())
		assert2.Equal("hash:mac", p.TypeName.Get())
		assert2.Len(p.Entries, 3)

		p = &res[2]
		assert2.Equal("baz", p.Name.Get())
		assert2.Equal("hash:ip", p.TypeName.Get())
		assert2.Len(p.Entries, 3)
	}

	m.AssertExpectations(t)
}

func TestConn_Add(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	m.On("Query", []byte{
		0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x66, 0x6f, 0x6f, 0x00, 0x34, 0x00, 0x08, 0x80, 0x18, 0x00, 0x07, 0x80, 0x0c, 0x00, 0x01, 0x80,
		0x08, 0x00, 0x01, 0x40, 0xc0, 0xa8, 0x01, 0x01, 0x08, 0x00, 0x09, 0x40, 0x00, 0x00, 0x00, 0x00,
		0x18, 0x00, 0x07, 0x80, 0x0c, 0x00, 0x01, 0x80, 0x08, 0x00, 0x01, 0x40, 0xc0, 0xa8, 0x01, 0x02,
		0x08, 0x00, 0x09, 0x40, 0x00, 0x00, 0x00, 0x01, 0x08, 0x00, 0x09, 0x40, 0x00, 0x00, 0x00, 0x00,
	}).Return([]netlink.Message{}, nil)

	c := Conn{Family: netfilter.ProtoIPv4, Conn: m}
	assert2.NoError(c.Add("foo",
		NewEntry(EntryIP(net.ParseIP("192.168.1.1"))),
		NewEntry(EntryIP(net.ParseIP("192.168.1.2")))))

	m.AssertExpectations(t)
}
