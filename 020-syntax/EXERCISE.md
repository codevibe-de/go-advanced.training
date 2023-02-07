# Übungen

## a) Variablen & Konstanten

Sie können für diese Übungen die Datei `020-syntax/a_variables-and-constants/main.go` nutzen.

### Variablen

Deklarieren Sie auf verschiedene Weisen Variablen.

Welche Zuweisungen von einer zu anderen gehen, welche nicht?

### Konstanten

Deklarieren Sie nicht-gruppierte und gruppierte Konstanten.

Was passiert, wenn diese gemischt in einem Ausdruck genutzt werden?

### Iota

Schreiben Sie eine Gruppe an Konstanten, welche die Werte `KB`, `MB` und `GB` deklariert.

Nutzen Sie hierfür `iota` zur automatischen Berechnung.

Tipp: mit `1 << 10` shiften Sie eine bestehende Zahl um 10 bits nach vorne

## b) Pointers

Nutzen Sie die Datei `020-syntax/b_pointers/main.go`.

### miniSort

Verändern Sie die Funktion `miniSort` und deren Aufruf derart, dass die gewünschte Ausgabe
erfolgt.

## c) Types

### Umwandlung cm in m

Deklarieren Sie eigene Typen für `centimeter` und `meter`. Wo muss die Deklaration erfolgen?

Schreiben Sie eine Funktion, die Zentimeter in Meter umrechnet.

Geben Sie das Ergebnis aus, sodass der gewünschte Output erfolgt.

## e) For Schleifen

Deklarieren Sie ein Array von Zahlen, z.B. `int` oder `float32`.

Schreiben Sie eine Funktion, welche das Array als Parameter annimmt und
die Summe aller Zahlen zurückgibt.

## f) If Abfragen

Schreiben Sie eine Funktion, welche für das Array von Zahlen (siehe oben)
nun sowohl den kleinsten wie auch den größten Wert zurückgibt.

## h) Strings und Runes

Für einen beliebigen String wie z.B. "Heute lerne ich Go" soll die Camel-Case 
Variante in der Form "HeuteLerneIchGo" berechnet werden.

Schreiben Sie die dafür notwendige Funktion.

## i) Structs

Sie möchten folgendes JSON mit Structs verarbeiten:

````json
{
  "produktId": "P-123",
  "NAME": "Pizza Salami",
  "Preis": {
    "betrag-je-stueck": 12.99,
    "währung": "EUR"
  }
}
````

Erstellen Sie die dafür nötigen Structs. Diese sollen nur englische Attribute verwenden
(also "productId" statt "produktId") und in Camel-Case notiert werden ("unitPrice" statt "betrag-je-stueck").

Legen Sie dann eine Struct Instanz mit den obigen Daten an.

## j) variadische Funktionen

### Statistiken

Ähnlich der Anzeige von Statistiken in Excel, wenn mehrere Zellen selektiert sind, benötigen wir
eine Funktion, die für eine beliebige Anzahl an `int` Werten die folgenden Statistiken berechnet:

- Anzahl Werte
- kleinster
- größter
- Summe
- Durchschnitt

## l) First Class Citizen Functions

Deklarieren Sie einen Funktionstyp mit folgender Signatur:

````go
funct(int, int) int
````

Implementieren Sie diese Funktion in mehreren Varianten:

* Addition
* Subtraktion
* Multiplikation

Rufen Sie nun ihre Funktion (auch gerne verschachtelt) auf und lassen
das Ergebnis ausgeben.

## ?) Slices

Implementieren Sie einen **Stack** von `string` Werten mithilfe eines Slice.

Dieser soll folgende Funktionen implementieren:

````go
func Push(stack []string, v string) {}
func Peek(stack []string) (string) {}
func Pop(stack []string) (string) {}
````