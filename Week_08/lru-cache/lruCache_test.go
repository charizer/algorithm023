package lru_cache

import "testing"

func TestConstructor(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {2=2, 1=1}
	want := 1
	got := lRUCache.Get(1)    // 返回 1
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {3=3, 1=1}
	want = -1
	got = lRUCache.Get(2)    // 返回 -1 (未找到)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	want = -1
	got = lRUCache.Get(1)    // 返回 -1 (未找到)
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	want = 3
	got = lRUCache.Get(3)    // 返回 3
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
	want = 4
	got = lRUCache.Get(4)    // 返回 4
	if got != want {
		t.Errorf("want:%d got:%d", want, got)
	}
}
