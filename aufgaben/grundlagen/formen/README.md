# Aufgaben zum Zeichnen von geometrischen Formen

In der Datei `formen.go` sind einige Funktionen vorgegeben.
Hier geht es darum, verschiedene geometrische Formen auf die Konsole zu zeichnen.
Die genauen Aufgabenstellungen sind bei den jeweiligen Funktionen beschrieben.

Die `main()`-Funktion ist leer.
Stattdessen gibt es Tests in der Datei `formen_test.go`.
Dort gibt es verschiedene Aufrufe der Funktionen und unten steht jeweils die
erwartete Ausgabe. Die Tests reichen vollkommen aus, um die Aufgaben zu bearbeiten.
Sie können alternativ die Funktionen aber auch aus der `main()` heraus aufrufen,
um sie zu testen.

## Erweiterte Aufgaben

### Doppelten Code vermeiden

Achten Sie beim Bearbeiten der Aufgaben darauf, so wenig doppelten Code wie möglich
zu schreiben. Dies wird auch das *DRY*-Prinzip genannt:
"**D**on't  **R**epeat **Y**ourself".

Doppelter Code macht auf Dauer die Wartung schwieriger und arbeitsintensiver und
er erhöht das Risiko von Fehlern.

Um doppelten Code zu vermeiden, sollten Sie versuchen,
Funktionen möglichst mit Hilfe anderer Funktionen zu schreiben.
Falls Sie im Nachhinein bemerken, dass Sie doppelten Code haben,
sollten Sie dies auch nachträglich noch korrigieren, indem Sie Funktionen umbauen.
Auch die Einführung weiterer Hilfsfunktionen kann nützlich sein.

### Benutzereingabe

Schreiben Sie in der `main()`-Funktion ein Programm, das den Benutzer fragt,
welche geometrische Form er mit welchen Zeichen in welcher Größe zeichnen will.
Das Programm soll anschließend die Aufgabe ausführen und sich beenden.
