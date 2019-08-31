package proximityhash

// stringQueue is a simple queue implementation that holds strings
type stringQueue struct {
	storage []string
	size    int
}

func newStringQueue() stringQueue {
	return stringQueue{size: 0}
}

func (s *stringQueue) enqueue(items ...string) {
	for _, item := range items {
		s.storage = append(s.storage, item)
		s.size = s.size + 1
	}
}

func (s *stringQueue) dequeue() (result string, ok bool) {
	if(s.isEmpty()) {
		return "", false
	}

	res := s.storage[0]
	s.storage = s.storage[1:]
	s.size = s.size - 1

	return res, true
}

func (s *stringQueue) isEmpty() bool {
	return s.size < 1
}
