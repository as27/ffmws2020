# Linked List

Wir implementieren eine einfache [Linked List](https://de.wikipedia.org/wiki/Liste_(Datenstruktur)) in Go. Die Liste kann dabei string Elemente aufnehmen.

Loop durch die Liste soll wie folgt möglich sein:

```Go
    for e := l.Front(); e != nil; e = e.Next() {
        // do something with e.Value
    }
```
[Quelle golang.org/pkg/container/list](https://golang.org/pkg/container/list/#pkg-index)

Es gibt dabei folgende Datentypen:

* Element
    * Value
    * next, prev = Pointer zu den jeweiligen Elementen
    
* Liste
    * currentNode
    * head

Element soll folgende Methoden implementieren:

* Next()*Element
* Prev()*Element

Die Liste sollte folgende Methoden implementieren:

* Front()*Element
* Len()int
* AddAt(int,string)
* Get(int)*Element
* PushFront(string): Element am Anfang einfügen
* PushBack(string) : Element am Ende einfügen
