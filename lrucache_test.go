package gostructures

import "testing"

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(5)

	if _, ok := cache.Fetch("key"); ok {
		t.Errorf("cache should not have key")
	}

	cache.Add("key", "value")
	cache.Add("key", "value")

	if value, ok := cache.Fetch("key"); !ok || value != "value" {
		t.Errorf("value should be %#v", value)
	}

	cache.Add("foo", "bar")
	cache.Add("bar", "baz")
	cache.Add("baz", "gaz")
	cache.Add("gar", "shaz")

	if cache.size != 5 {
		t.Errorf("size should be 5, actual: %#v", cache.size)
	}

	// "key" is the LRU so should be kicked out
	if !cache.HasKey("key") {
		t.Errorf("cache should have key")
	}

	cache.Add("shaz", "blaz")

	if cache.HasKey("key") {
		t.Errorf("cache should no longer have key")
	}

	if cache.size != 5 {
		t.Errorf("size should still be 5, actual: %#v", cache.size)
	}

	// next up "foo" should get kicked out
	// but we'll use it again
	cache.Add("foo", "barbie")

	if value, ok := cache.Fetch("foo"); !ok || value != "barbie" {
		t.Errorf("value should be %#v", value)
	}

	cache.Add("blaz", "berg")

	if !cache.HasKey("foo") {
		t.Errorf("foo should still be around")
	}

	if cache.HasKey("bar") {
		t.Errorf("bar should be gone")
	}
}
