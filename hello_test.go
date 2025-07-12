package main

import (
	"os"
	"os/exec"
	"testing"
)

// TestHelloOutput prüft, ob das Programm kompilierbar ist und "Hallo, Welt!" ausgibt
func TestHelloOutput(t *testing.T) {
	// Kompiliere das Programm
	cmd := exec.Command("go", "build", "-o", "hello.test")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Kompilierung fehlgeschlagen: %v", err)
	}
	defer os.Remove("hello.test") // Aufräumen nach dem Test

	// Führe das kompilierte Programm aus
	cmd = exec.Command("./hello.test")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Ausführung fehlgeschlagen: %v", err)
	}

	// Überprüfe die Ausgabe
	expected := "Hallo, Welt!\n"
	if string(output) != expected {
		t.Fatalf("Erwartete Ausgabe: %q, aber erhalten: %q", expected, string(output))
	}
}
