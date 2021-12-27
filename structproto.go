package structproto

func Prototypify(target interface{}, option *StructProtoResolveOption) (*Struct, error) {
	if target == nil {
		panic("specified argument 'target' cannot be nil")
	}

	r := NewStructProtoResolver(option)
	prototype, err := r.Resolve(target)
	if err != nil {
		return nil, err
	}
	return prototype, nil
}
