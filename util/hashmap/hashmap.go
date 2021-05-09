package hashmap

import (
	"reflect"
)

type HashMap struct {
	m map[interface{}]interface{}
}

//HashMap()
func New() *HashMap {
	return &HashMap{m: make(map[interface{}]interface{})}
}

//HashMap(int initialCapacity)
func NewWithInitialCapacity(initialCapacity int) *HashMap {
	return &HashMap{m: make(map[interface{}]interface{}, initialCapacity)}
}

//clear()
func (m *HashMap) Clear() {
	m.m = make(map[interface{}]interface{})
}

//compute(K key, BiFunction<? super K,? super V,? extends V> remappingFunction)
//computeIfAbsent(K key, Function<? super K,? extends V> mappingFunction)
//computeIfPresent(K key, BiFunction<? super K,? super V,? extends V> remappingFunction)

//containsKey(Object key)
func (m *HashMap) ContainsKey(key interface{}) bool {
	_, ok := m.m[key]
	return ok
}

//containsValue(Object value)
func (m *HashMap) ContainsValue(value interface{}) bool {
	for _, v := range m.m {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

//entrySet()
//forEach(BiConsumer<? super K,? super V> action)

//get(Object key)
func (m *HashMap) Get(key interface{}) interface{} {
	value, _ := m.m[key]
	return value
}

//getOrDefault(Object key, V defaultValue)
func (m *HashMap) GetOrDefault(key, V interface{}) interface{} {
	value, ok := m.m[key]
	if ok {
		return value
	}
	return V
}

//isEmpty()
func (m *HashMap) IsEmpty() bool {
	if len(m.m) == 0 {
		return true
	}
	return false
}

//keySet()

//merge(K key, V value, BiFunction<? super V,? super V,? extends V> remappingFunction)

//put(K key, V value)
func (m *HashMap) Put(K, V interface{}) interface{} {
	m.m[K] = V
	return V
}

//putAll(Map<? extends K,? extends V> m)
//putIfAbsent(K key, V value)
func (m *HashMap) PutIfAbsent(K, V interface{}) interface{} {
	value, ok := m.m[K]
	if ok {
		return value
	}
	m.m[K] = V
	return nil
}

//remove(Object key)
func (m *HashMap) Remove(key interface{}) interface{} {
	value, ok := m.m[key]
	if ok {
		delete(m.m, key)
		return value
	}
	return nil
}

//remove(Object key, Object value)
func (m *HashMap) RemoveWithKeyValue(K, V interface{}) bool {
	value, ok := m.m[K]
	if ok && reflect.DeepEqual(V, value) {
		delete(m.m, K)
		return true
	}
	return false
}

//replace(K key, V value)
func (m *HashMap) Replace(K, V interface{}) interface{} {
	value, ok := m.m[K]
	if ok {
		m.m[K] = V
		return value
	}
	return nil
}

//replace(K key, V oldValue, V newValue)
func (m *HashMap) ReplaceWithValue(K, oldValue, newValue interface{}) bool {
	value, ok := m.m[K]
	if ok && reflect.DeepEqual(value, oldValue) {
		m.m[K] = newValue
		return true
	}
	return false
}

//replaceAll(BiFunction<? super K,? super V,? extends V> function)

func (m *HashMap) Size() int {
	return len(m.m)
}

//values()
