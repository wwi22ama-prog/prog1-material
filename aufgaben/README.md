# Übungsaufgaben zum Programmieren

Hier finden sich verschiedene Unterordner mit Aufgaben zu unterschiedlichen Themen.
Diese Themen-Ordner enthalten wiederum Unterordner mit den einzelnen Aufgaben.

## Aufbau der Aufgaben

In jedem Aufgaben-Ordner gibt es eine Reihe von Dateien:

* Eine oder mehrere Quelldateien, in denen Code vorgegeben und die Aufgabe
  beschrieben ist. In diesen Dateien sollen Sie arbeiten.
  Sie können weitere eigene Dateien und/oder Funktionen anlegen.
* Eine oder mehrere Testdateien, deren Name auf `_test.go` endet.
  Diese Dateien enthalten die Tests für die Aufgaben.
  Dies sind Funktionen, die das Verhalten der Funktionen aus den Aufgaben prüfen sollen.
  Sie können die Tests über Ihre IDE oder auf der Kommandozeile mit dem Befehl
  `go test` im jeweiligen Aufagben-Ordner ausführen.
  Man kann auch einzelne Tests gezielt ausführen oder die Ausführung genauer steuern,
  geben Sie dazu am Besten auf der Konsole `go help test` ein.

## Technisches

Damit das ganze Framework aus Go und seinen Tests reibungslos funktioniert,
müssen einige Voraussetzungen erfüllt sein.
Alle vorgegebenen Ordner mit Code-Beispielen oder Aufgaben enthalten
Dateien `go.mod`. Darin ist jeweils ein eindeutiger Name für das `Modul` enthalten,
das hier definiert wird.

Die Namen in diesem Repo sind gültige Web-Adressen.
Wenn Sie eigene Module schreiben, die Sie nicht veröffentlichen, können Sie hier
einen völlig beliebigen eindeutigen Namen wählen.
Sie können die `go.mod` von Hand anlegen bzw. kopieren und abändern,
oder sie verwenden dafür den Befehl `go mod init <NAME>` auf der Konsole.

Außerdem gibt es im Hauptverzeichnis eine Datei `go.work`, in der alle
Unterverzeichnisse aufgelistet sind, in denen Module definiert werden.
Wenn Sie ein neues Modul schreiben, müssen Sie sein Verzeichnis dort hinzufügen,
entweder von Hand oder mit dem Befehl `go work`.
