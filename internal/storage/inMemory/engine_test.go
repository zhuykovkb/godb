package inMemory_test

import (
	"github.com/stretchr/testify/assert"
	"goconcurrency/internal/storage/inMemory"
	"testing"
)

func TestEngine_SetGet(t *testing.T) {
	en := inMemory.NewEngine()

	en.Set("key1", "value1")
	en.Set("key2", "value2")

	val, ok := en.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)

	val, ok = en.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "value2", val)

	val, ok = en.Get("key3")
	assert.False(t, ok)
	assert.Empty(t, val)
}

func TestEngine_UpdateValue(t *testing.T) {
	en := inMemory.NewEngine()

	en.Set("key", "initial_value")
	val, ok := en.Get("key")
	assert.True(t, ok)
	assert.Equal(t, "initial_value", val)

	en.Set("key", "updated_value")
	val, ok = en.Get("key")
	assert.True(t, ok)
	assert.Equal(t, "updated_value", val)
}

func TestEngine_Del(t *testing.T) {
	en := inMemory.NewEngine()

	en.Set("key1", "value1")
	en.Del("key1")

	val, ok := en.Get("key1")
	assert.False(t, ok)
	assert.Empty(t, val)

	en.Del("key1")
}

func TestEngine_CaseSensitivity(t *testing.T) {
	en := inMemory.NewEngine()

	en.Set("Key", "value1")
	en.Set("key", "value2")

	val, ok := en.Get("Key")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)

	val, ok = en.Get("key")
	assert.True(t, ok)
	assert.Equal(t, "value2", val)
}

func TestEngine_EmptyKeyAndValue(t *testing.T) {
	en := inMemory.NewEngine()

	en.Set("", "empty_key_value")
	val, ok := en.Get("")
	assert.True(t, ok)
	assert.Equal(t, "empty_key_value", val)

	en.Set("empty_value_key", "")
	val, ok = en.Get("empty_value_key")
	assert.True(t, ok)
	assert.Equal(t, "", val)
}

func TestEngine_LargeValues(t *testing.T) {
	en := inMemory.NewEngine()

	largeKey := "key" + string(make([]byte, 1000))
	largeValue := "value" + string(make([]byte, 1000))

	en.Set(largeKey, largeValue)
	val, ok := en.Get(largeKey)
	assert.True(t, ok)
	assert.Equal(t, largeValue, val)
}

func TestEngine_MultipleOperations(t *testing.T) {
	en := inMemory.NewEngine()

	en.Set("key1", "value1")
	en.Set("key2", "value2")
	en.Set("key3", "value3")
	en.Del("key2")

	_, ok := en.Get("key2")
	assert.False(t, ok)

	val, ok := en.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)

	val, ok = en.Get("key3")
	assert.True(t, ok)
	assert.Equal(t, "value3", val)
}
