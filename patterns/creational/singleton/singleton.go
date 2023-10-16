package singleton

type singleton struct {
	chars []byte
}

var singletonOnce *singleton

func (s *singleton) AddChar(r byte) {
	s.chars = append(s.chars, r)
}

func ReadChars() []byte {
	if singletonOnce == nil {
		singletonOnce = new(singleton)
	}

	return singletonOnce.chars
}
