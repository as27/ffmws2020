# Typen

Go ist eine Sprache mit einer [strengen Typisierung](https://de.wikipedia.org/wiki/Starke_Typisierung). Das bedeutet, dass der Typ einer Variablen im Code mit angegeben werden muss. Bei Sprachen mit einer schwachen Typisierung wird der Typ der Variablen zur Laufzeit des Programmes erst ermittelt. 



Anders ausgedrückt, wenn einmal einer Variablen ein Typ zugewiesen wird, so kann dieser auch nicht mehr verändert werden.

In JavaScript (oder auch Python) wird dies nicht so streng gehandhabt. Dies führt in der Praxis immer wieder zu sehr schwer auffindbaren Fehlern. Das ist auch der Grund, warum sich in den letzten Jahren immer mehr [TypeScript](https://de.wikipedia.org/wiki/TypeScript) durchsetzt. TypeScript ist praktisch JavaScript mit Typen.

Der Vorteil der starken Typisierung ist 

* eine erhöhte Lesbarkeit, da bei der Variablendeklaration zusätlich auch der Typ mitgegeben wird
* stabilerer Code, da der Compiler ein Mischen von Typen nicht zulässt. 


Sollte in Go einer Variablen mit dem Typ `string` (also Text) ein `integer` (also Ganzzahl) zugewiesen werden, so lässt sich das Programm gar nicht kompilieren. Der Compiler meldet hier frühzeitig einen Fehler (https://play.golang.org/p/9OdsHARtZE).

```Go
var a string
a = 1
```
```
cannot use 1 (type int) as type string in assignment
```

Eine gute Übersicht über alle Basistypen bietet die Go Tour: https://go-tour-de.appspot.com/basics/11

In der [Dokumentation der Packetes `builtin`](https://golang.org/pkg/builtin/) werden die einzelnen Typen auch genau beschrieben.

Für die erste Verwendung im Workshop beschränken wir uns auf die Typen:

* bool (kann true oder false sein)
* string (representiert einen Text)
* int (ist eine Ganzzahl)
* byte (ein Byte, wird z.B. verwendet, wenn eine Datein eingelesen wird)

## Deklaration

Die Deklaration von Basistypen in Go sieht wie folgt aus:


```Go
var myString string
myString = "mein Text"
var myInteger int
myInteger = 4
// Inklusive Initialisierung
var anotherInteger int = 4
```

[Beispiele aus der Go Tour](https://go-tour-de.appspot.com/basics/8)

Es gibt auch die Möglichkeit einer kurzen Variablendeklaration. Allgemein wird diese Deklaration in den meisten Projekten bevorzugt verwendet. Die direkte Lesbarkeit nimmt dadurch ein wenig ab. Bei der Verwendung eines Editors wie VS Code ist dies nicht so schlimm, da dieser einem auch den Typen der Variablen anzeigt.

```Go
myInteger := 4
```

Jeder Typ besitzt auch eine Funktion, welche diesen Wert zurück gibt. Hierfür wird dem Typen `()` angehängt. Diese Funktion wandelt, solange es möglich ist, den Wert in den jeweiligen Typen um.

```Go
myInteger := int(4)
myString := string("mein Text")
```

## Eigene Typen

Ein weiterer wichtiger Bestandteil ist die Definition von eigenen Typen. Hierdurch können die Basistypen erweitert werden. In folgendem Beispiel werden Typen für die Temparatur definiert. 

```Go
type Celcius int
type Farenheit int

func main() {
	tc := Celcius(24)
	tf := Farenheit(72)
	fmt.Println("tc: ", tc)
	fmt.Println("tf: ", tf)
}
```
https://play.golang.org/p/hyKcUrhym7

Obwohl beide Typen jeweils ein `int` sind, ist es in Go nicht möglich diese beiden Typen zu mischen:

```Go
type Celcius int
type Farenheit int

func main() {
	tc := Celcius(24)
	tf := Farenheit(72)
	t := tc + tf
}
```
https://play.golang.org/p/D_-RxDM2HI

Die Fehlermeldung hier:
```
invalid operation: tc + tf (mismatched types Celcius and Farenheit)
```

In JavaScript wäre hier eine Addition ohne Fehler möglich gewesen. D.h. bei dem Design der Programme müssen fachliche Typen identifiziert werden, welche dann auch im Code definiert werden. Bei einer guten Aufteilung werden dann Fehler bereits frühzeitig erkannt.

## Zusammenfassung

Go ist eine Sprache mit strenger Typisierung. Die einzelnen Typen können nicht vermischt werden. Zu den Basistypen können eigene fachlich definierte Typen erzeugt werden.

Dies ist eine erste einfache Einführung zu den Typen. In den kommenden Kapiteln werden wir immer wieder auf die Typen zurück kommen und weitere Definitionen kennen lernen.


