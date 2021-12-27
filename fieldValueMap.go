package structproto

var _ FieldValueCollectionIterator = new(FieldValueMap)

type FieldValueMap map[string]interface{}

func (values FieldValueMap) Iterate() <-chan FieldValueEntry {
	c := make(chan FieldValueEntry, 1)
	go func() {
		for k, v := range values {
			c <- FieldValueEntry{k, v}
		}
		close(c)
	}()
	return c
}
