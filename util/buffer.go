package util

type StrArrBuffer struct {
	buf []string
}

func (s *StrArrBuffer) Write(content string) {
	s.buf = append(s.buf, content)
}
func (s *StrArrBuffer) GetArr() []string {
	return s.buf
}
