package hashtable

type bucket = map[string]any

type hashTable struct {
	table map[int32]bucket
}

func New() *hashTable {
	return &hashTable{
		table: make(map[int32]map[string]any),
	}
}

func (h *hashTable) Get(key string) (any, bool) {
	if key == "" {
		return nil, false
	}
	position := h.djb2Hash(key)
	if bucket, ok := h.table[position]; ok {
		val, exists := bucket[key]
		return val, exists
	} else {
		return nil, false
	}
}

func (h *hashTable) Put(key string, value any) bool {
	if key != "" && value != nil {
		position := h.djb2Hash(key)
		if h.table[position] == nil {
			h.table[position] = make(map[string]any)
		}
		h.table[position][key] = value
		return true
	}
	return false
}

func (h *hashTable) Delete(key string) bool {
	if key == "" {
		return false
	}
	position := h.djb2Hash(key)
	if bucket, ok := h.table[position]; ok {
		_, exists := bucket[key]
		if exists {
			delete(h.table[position], key)
		}
		return exists
	}

	return false
}

func (h *hashTable) djb2Hash(key string) int32 {
	var hash int32 = 5381
	var spreader int32 = 33

	for _, char := range key {
		hash = (hash * spreader) + int32(char)
	}

	if hash < 0 {
		hash = -hash
	}

	return hash % 1013 // tamnaho total de 1.000
}
func (h *hashTable) loseloseHash(key string) int32 {
	var hash int32
	for _, char := range key {
		hash += char
	}

	return hash % 37
}
