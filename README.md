# PCP_GO 

## Aufgaben die gelöst wurden
- Aufgaben aus der Vorlesung:
  - Fibonacci
  - Collatz
  - Parallelität (go-routines, channels)
- Aufgaben aus den Themenpunkten:
  - Maps & Slices
  - Structural Typing
  - Defer/Panic/Recover
  - go-routines & channels

## 1. Installation 
### **Installation MacOS**
* Lade das macOS-Paket (`.pkg`) von der offiziellen Website herunter: [golang.org/dl](https://golang.org/dl/)
* Öffne die Datei und folge dem Installationsassistenten.
* oder mit Homebrew:
```bash
brew install go
```
* Verifikation: Öffne das Terminal und tippe:
```bash
go version
```

### Installation Windows
1. **Per Installer:**
* Lade den Windows-Installer (`.msi`) herunter: [golang.org/dl](https://golang.org/dl/)
* Starte den Installer und folge den Anweisungen (der Installer setzt automatisch die Umgebungsvariablen wie `PATH`).

2. **Verifikation:**
   Öffne die Eingabeaufforderung (`cmd`) oder die PowerShell und tippe:
```cmd
go version

```

---

## 2. Ausführung im Terminal

Navigiere auf macOS oder Windows (Terminal / PowerShell) in das jeweilige Verzeichnis des gewünschten Themas und führe den Code direkt aus.

### Befehl zum direkten Ausführen:

```bash
# Beispiel für Maps & Slices
cd code/fib 
go run main.go

# Beispiel für Structural Typing
cd code/structural-typing
go run main.go

# Beispiel für Defer/Panic/Recover
cd code/defer_panic_recover
go run main.go
```

### Befehl zum Kompilieren in eine native Binärdatei (Executable):

Wenn du eine eigenständige, ausführbare Datei generieren möchtest:

```bash
go build main.go
```

* Unter **macOS** entsteht eine Unix-Binärdatei `./main`.
* Unter **Windows** entsteht eine ausführbare Datei `main.exe`.