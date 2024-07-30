package xmlstream

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push("a", "b", "c")
	if s.Len() != 3 && s.Peek() != "c" {
		t.Fatalf("expected depth 3 and emerging \"c\", got %d and %q", s.Len(), s.Peek())
	}
	if s.At(-1) != "c" {
		t.Fatalf("expected element \"c\" at -1, got %q", s.At(-1))
	}
	if s.At(-2) != "b" {
		t.Fatalf("expected element \"b\" at -1, got %q", s.At(-1))
	}
	if s.At(-3) != "a" {
		t.Fatalf("expected element \"a\" at -1, got %q", s.At(-1))
	}
	s.Pop()
	if s.Len() != 2 && s.Peek() != "b" {
		t.Fatalf("expected depth 2 and emerging \"b\", got %d and %q", s.Len(), s.Peek())
	}
	s.Pop()
	if s.Len() != 1 && s.Peek() != "2" {
		t.Fatalf("expected depth 1 and emerging \"a\", got %d and %q", s.Len(), s.Peek())
	}
	s.Pop()
	if s.Len() != 0 && s.Peek() != "" {
		t.Fatalf("expected depth 0 and emerging \"\", got %d and %q", s.Len(), s.Peek())
	}
	if !s.Empty() {
		t.Fatalf("expected empty, got depth of %d", s.Len())
	}
}
