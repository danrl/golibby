package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpsert(t *testing.T) {
	{
		hm := HashMap{}
		assert.Equal(t, 1<<16, len(hm.data))
	}
	{
		hm := HashMap{}

		hm.Upsert("foo", "bar")
		assert.Equal(t, 1, len(hm.data[0x1bc5]))
		assert.Equal(t, "foo", hm.data[0x1bc5][0].key)
		assert.Equal(t, "bar", hm.data[0x1bc5][0].value)

		hm.Upsert("foo", "updated")
		assert.Equal(t, 1, len(hm.data[0x1bc5]))
		assert.Equal(t, "foo", hm.data[0x1bc5][0].key)
		assert.Equal(t, "updated", hm.data[0x1bc5][0].value)

		hm.Upsert("Zero-byte", "collision")
		assert.Equal(t, 2, len(hm.data[0x1bc5]))
		assert.Equal(t, "Zero-byte", hm.data[0x1bc5][1].key)
		assert.Equal(t, "collision", hm.data[0x1bc5][1].value)
	}
}

func TestDelete(t *testing.T) {
	{
		hm := HashMap{}

		err := hm.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
	// delete single
	{
		hm := HashMap{}

		hm.Upsert("foo", "bar")
		err := hm.Delete("foo")
		assert.Equal(t, nil, err)
		assert.Equal(t, 0, len(hm.data[0x1bc5]))
	}
	// delete first
	{
		hm := HashMap{}

		hm.Upsert("foo", "bar")
		hm.Upsert("Zero-byte", "collision")
		hm.Upsert("Water of hydration", "H2O")

		err := hm.Delete("foo")
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, len(hm.data[0x1bc5]))
		assert.Equal(t, "Zero-byte", hm.data[0x1bc5][0].key)
		assert.Equal(t, "collision", hm.data[0x1bc5][0].value)
		assert.Equal(t, "Water of hydration", hm.data[0x1bc5][1].key)
		assert.Equal(t, "H2O", hm.data[0x1bc5][1].value)
	}
	// delete middle
	{
		hm := HashMap{}

		hm.Upsert("foo", "bar")
		hm.Upsert("Zero-byte", "collision")
		hm.Upsert("Water of hydration", "H2O")

		err := hm.Delete("Zero-byte")
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, len(hm.data[0x1bc5]))
		assert.Equal(t, "foo", hm.data[0x1bc5][0].key)
		assert.Equal(t, "bar", hm.data[0x1bc5][0].value)
		assert.Equal(t, "Water of hydration", hm.data[0x1bc5][1].key)
		assert.Equal(t, "H2O", hm.data[0x1bc5][1].value)
	}
	// delete last
	{
		hm := HashMap{}

		hm.Upsert("foo", "bar")
		hm.Upsert("Zero-byte", "collision")
		hm.Upsert("Water of hydration", "H2O")

		err := hm.Delete("Water of hydration")
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, len(hm.data[0x1bc5]))
		assert.Equal(t, "foo", hm.data[0x1bc5][0].key)
		assert.Equal(t, "bar", hm.data[0x1bc5][0].value)
		assert.Equal(t, "Zero-byte", hm.data[0x1bc5][1].key)
		assert.Equal(t, "collision", hm.data[0x1bc5][1].value)
	}
}

func TestValue(t *testing.T) {
	{
		hm := HashMap{}

		_, err := hm.Value("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		hm := HashMap{}

		hm.Upsert("foo", "bar")
		value, err := hm.Value("foo")
		assert.Equal(t, "bar", value)
		assert.Equal(t, nil, err)

		hm.Upsert("Zero-byte", "collision")
		value, err = hm.Value("Zero-byte")
		assert.Equal(t, "collision", value)
		assert.Equal(t, nil, err)
	}
}
