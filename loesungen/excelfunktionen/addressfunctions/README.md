# Package addressfunctions

Dieses Package enthält Funktionen zur Analyse von Adressen der Tabellenkalkulation.
Es gibt Funktionen, die Adressen zu Zeilen- und Spaltennummern auswerten können,
sowie Funktionen, die für eine gegebene Adresse bestimmen können, ob diese ein
Zell- oder ein Bereichs-Bezug ist.

## Voraussetzungen: Die String-Helfer aus dem Package misc

Die Funktionen in `stringhelpers.go` aus dem Package `misc` könnten hier hilfreich sein.
Diese sollten Sie also zuerst implementieren oder zwischendurch darauf zurückkommen.

**Anwendung der Funktionen:** Um z.B. die Funktion `IsDigit()` aus dem Package `misc`
in einer der Dateien hier zu verwenden, muss sie mit `misc.IsDigit(...)` aufgerufen
werden. Außerdem muss das Package mit `"loesungen/excelfunktionen/misc"` importiert werden.
