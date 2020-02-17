package list

type Element struct {
	Value string
	next  *Element
	prev  *Element
}

func (e *Element) Next() *Element { return e }

func (e *Element) Prev() *Element { return e }

type List struct {
	currentElement *Element
	head           *Element
}

func (l *List) Front() *Element { return nil }

func (l *List) AddAt(i int, v string) error { return nil }

func (l *List) Get(i int) (*Element, error) { return nil, nil }

func (l *List) Len() int { return 0 }
