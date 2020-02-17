# Go Test

Go unterstützt von natur aus automatische Tests. Diese Tests machen wir uns beim Workshop bei den Übungen zu nutze. Innerhalb der Übungen sind Tests vorbereitet. Diese Tests prüfen dann, ob die Aufgaben auch korrekt gelöst wurden. Auf diese Weise kann jeder selber kontrollieren, ob seine Lösung auch korrekt ist.

Das ist auch der Grund, warum vor allen anderen Elementen die Tests gleich als erstes vorgestellt werden.

## Die Testdatei

Eine Testdatei wird immer zu einer Quelldatei erzeugt. Dafür wird dem Dateinamen `_test` angehängt.

Tests für `main.go` liegen immer in der Datei `main_test.go` oder bei `server.go` in `server_test.go`.

Auf diese Weise werden die zugehörigen Testdateien im Verzeichnig gleich bei der Go Datei angezeigt.

## Test über die Commandozeile

Die Go Tests werden meistens über einen Terminal gestartet. Der Befehl hierfür ist `go test`. Damit werden alle tests des aktuellen Verzeichnises getestet.

Möchte man alle Tests eines Projektes ausführen so verwendet man.

```
go test ./...
```

## Test in Visual Studio Code

VS Code unterstützt die Tests in Go. Sobald die Go Erweiterung installiert ist, erkennt VS Code automatisch die Testdateien und die dazugehörigen Tests. Oberhalb der Funktionen werden kleine Links eingeblendet, über welche die einzelnen Tests direkt aus VS Code aus gestartet werden können. 

# gofmt und goimports

Teil des Go Toolings sind die Programme zur Formatierung des Go Codes. Go wurde mit der Intention der Lebarkeit erstellt. Es wurden dafür klare Regeln erstellt, wie guter Go Code formatiert sein muss. Damit dies überall einheitlich ist, wurde ein eigenes Tool zur Formatierung des Codes erstellt. Durch den Befehl `go fmt meinegodatei.go` wird die Datei automatisch korrekt formatiert.

Eine Erweiterung zur Formatierung stellt das Programm goimports dar. Dieses Programm fügt automatisch zur Formatierung noch die `import` Anweisungen für die Packete ein.

Diese Tools sind auch im VS Code eingebunden. Jedes mal beim Speichern einer Datei wird der Code formatiert und die `import` Anweisungen gesetzt.

# govet

Für idiomatischen Go Code gibt es noch weitere Empfehlungen. Diese können mit `go vet` aufgerufen werden. VS Code führt auch diesen Befehl nach jedem Speichern aus und markiert die Zeilen, welche Anmerkungen hierzu besitzen.

# go run

Mit dem Befehl `go run datei.go` kann ein Go Programm direkt ausgeführt werden. 