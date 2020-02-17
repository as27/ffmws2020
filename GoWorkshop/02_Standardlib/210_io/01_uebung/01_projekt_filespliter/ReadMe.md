# Projekt Filespliter

Vor vielen vielen Jahren, als Daten noch auf Disketten gespeichert wurden, gab es Programme, welche große Dateien in kleinere Einzeldateien auteilten. 

Diese Programme werden _Filespliter_ genannt. Ziel dieser Übung ist es genau so einen _Filespliter_ in Go zu implementieren. 

## Ziele

Mit diesem Projekt sollst Du verschiedene Fähigkeiten erweben bzw. üben.

* Verwenden der Dokumentation der Standar library
* Verwenden des io.Reader und io.Writer
* Erstellen eigener Tests

## Aufgabenstellung

Aus einer beliebigen Datei (jpg, png, txt, doc, zip, ...) werden meherere kleine Dateien (chunks) erzeugt. 

Als Input Parameter bekommt das Programm eine Datei (als Dateiname), die Größe der Chunks in Bytes und einen Output Folder für die Chunks. Die Eingangswerte können zuerst als Varaiblen deklariert werden. Nachdem der Splitter funktioniert strukturieren wir das Programm um. Dafür wird dann das [flag Packet](https://golang.org/pkg/flag/) zu verwendet.

Die split Funktionalitäten werden innerhalb eines eigenen Packetes gebündelt. Dieses Packet verwendet als Input immer den io.Reader bzw. auch einen io.Writer. Die Verwendung dieser beiden Interfaces wird als idiomatisches Go bezeichnet. 

Das Packet Split gibt bereits eine Struktur vor. Zusätzlich ist die Schnittstelle zum Teil bereits im Code Dokumentiert, sodass bereits eine Dokumentation erzeugt wird.

Für die Funktion Split() ist bereits auch schon ein Test vorgegeben und für Merge() ein Beispiel angelegt. 

Folgende Aufgaben sind nun zu erfüllen:

1) Vervollständige die Split() Funktion, so dass die Tests durchlaufen.
2) Erstelle einen sinnvollen Unit Test für die Funktion Merge().
3) Ergänze den Code von Merge() sodass Deine Tests grün werden (=PASS).
4) Erweitere die main() Funktion, damit diese eine Datei in mehrere Dateien aufteilt.
5) Erweitere die Datei main.go so, dass ein CLI (Command Line Interface) vorhanden ist. Hierüber kann Chunksize, Inputdatei und Outputfolder festgelegt werden. Dabei soll das [flag Packet](https://golang.org/pkg/flag/) verwendet werden.
6) Erweitere das Tool um ein Flag merge, welches die chunks eines Ordners wieder zusammenfügt.
