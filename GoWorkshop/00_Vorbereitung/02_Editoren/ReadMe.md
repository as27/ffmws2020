# Editor

## Visual Studio Code

Für den Workshop empfehle ich [Visual Studio Code](https://code.visualstudio.com/). Dieser ist für alle Systeme verfügbar und die Installation ist unkompliziert. Für Go gibt es eine sehr gute Erweiterung. Deshalb wird im gesamten Workshop Visual Studio Code verwendet.

### Installation

Damit die Go Erweiterung korrekt installiert werden kann muss Go bereits installiert sein, da die in der Erweiterung vorhandenen Tools in Go geschrieben sind.

* Lade den Installer / Packet für Dein System herunter: https://code.visualstudio.com/
* Installiere Visual Studio Code
* Gehe zu Erweiterungen und suche nach _Go_ und wähle das [Plugin von _lukehoban_](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go) aus. Installiere die Erweiterung

### Go Tools für die Erweiterung

Damit die Tools auch funktionieren müssen diese manuell installiert werden. Dafür ist es notwendig, dass Go bereits korrekt installiert ist. Dann kann mit dem Befehl `go get` das Tool einfach installiert werden.

Am Einfachsten ist es die folgenden Anweisungen einfach per Copy&Paste in den Terminal einzufügen. 

```
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/acroca/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/cweill/gotests/...
go get -u -v golang.org/x/tools/cmd/godoc
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/josharian/impl
```
Die Tools werden dann geladen und installiert. Komplette Anleitung findet sich im [vscode-go wiki](https://github.com/Microsoft/vscode-go/wiki/Go-tools-that-the-Go-extension-depends-on).

## Andere Editoren

* [JetBrains Gogland](https://www.jetbrains.com/go/)
* [vim-go Erweiterung für vim](https://github.com/fatih/vim-go)