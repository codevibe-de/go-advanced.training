# Übungen zum Thema "Concurrency"

## a) Channels: Lesen und Schreiben

Starten Sie eine Goroutine, die in einen Channel einen Wert schreibt (z.B. ein `int`).

Dieser Wert soll in der `main()` Funktion ausgelesen und ausgegeben werden.

## b) Gemeinsamer Start

Bei dieser Übung geht es darum, dass wir drei Goroutinen starten, die eine unbekannte Dauer brauchen, bis sie
einsatzbereit sind. Sobald wir uns sicher sind, dass alle drei bereit sind, sollen sie gleichzeitig etwas tun.

Hierfür starten Sie drei Goroutinen, die anfangs jeweils die vorgegebene Funktion `startup()` aufrufen,
um das unterschiedlich lange initialisieren (bereitmachen) zu simulieren.

Dann: Sobald jede Goroutine bereit ist, soll sie auf ein gemeinsames Startsignal (mittels Channel) warten und dann
die aktuelle Zeit ausgeben.

Wie wissen wir, wann jede Goroutine bereit ist? Mögliche Optionen, von einfach nach kompliziert:

1. in `main()` einfach viele Sekunden warten
2. Oder man löst es mit einer `sync.WaitGroup`
3. Oder mit einem zusätzlichen Kanal

In jedem Fall schreiben wir dann in den Startsignal-Channel Werte, um jede Goroutine nahezu gleichzeitig
loslaufen zu lassen.

## c) Entweder Oder

Starten Sie eine Goroutine, die alle 800 ms einen `int` Wert in einen Kanal schreibt.

Starten Sie eine andere Goroutine, die alle 1100 ms einen `string` in einen anderen Kanal schreibt.

In der `main()` Methode nutzen Sie ein `select`, um entweder auf die eine oder auf die andere Nachricht zu reagieren.

Ihre Anwendung wird endlos laufen, Sie können diese u.a. mit CTRL+C abbrechen.

## d) Pizza

Wir backen Pizzen! 🍕🍕🍕

Schauen Sie sich den Source-Code in `070_concurrency/d_pizza/main.go` an. Dort werden Pizzen und ihre Zutaten
definiert. Es gibt auch einen Ofen. Sowohl die Zutaten als auch der Ofen müssen präpariert werden, dies benötigt
jeweils Zeit. Das Backen der Pizzen ebenfalls.

Ein rein sequentieller Ablauf ist vorgegeben.

Aber in Ihrer Pizzeria haben wir drei tüchtige Helfer, nicht nur einen.

Reduzieren Sie die Gesamtzeit so weit wie möglich, indem sie mit insgesamt 3 laufenden Prozessen arbeiten. Die
bestehenden Datentypen können erweitert werden, falls Sie noch Methoden oder Interfaces einbauen wollen.

ACHTUNG, diese Übung ist wesentlich komplizierter als anfangs gedacht. Erste Hilfe finden Sie in der
Konzeptdatei [pizza-solution-concept.md](d_pizza/pizza-solution-concept.md).

Im `solution` Branch finden Sie eine mögliche Lösung.

ITERATIVES VORGEHEN: Sie können die Übung in steigender Komplexität angehen, z.B.:

- erst einmal nur Zutaten über Worker-Goroutinen vorbereiten lassen, sobald diese fertig sind die Pizzen in main()
  backen
- dann auch den Ofen über einen Worker vorheizen lassen
- dann das Backen der Pizzen auch über Goroutinen laufen lassen (nach Vorbereitung **aller** Zutaten)
- dann eine Pizza per Goroutine backen lassen, sobald nur deren Zutaten fertig sind

Diese Übung soll Spaß machen! Wenn es zu kompliziert wird, dann bitte die Zielsetzung vereinfachen! 🚀