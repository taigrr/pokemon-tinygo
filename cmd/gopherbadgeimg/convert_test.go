package main

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"testing"
)

func TestImgToBytes_AllBlack(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 8))
	for x := 0; x < 16; x++ {
		for y := 0; y < 8; y++ {
			img.Set(x, y, color.Black)
		}
	}
	var src image.Image = img
	disableDithering = true
	defer func() { disableDithering = false }()

	result := ImgToBytes(16, 8, &src)
	if len(result) != 16 {
		t.Fatalf("expected 16 bytes, got %d", len(result))
	}
	// All black pixels should produce all-set bits
	for i, b := range result {
		if b != 0xFF {
			t.Errorf("byte %d: expected 0xFF, got 0x%02X", i, b)
		}
	}
}

func TestImgToBytes_AllWhite(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 8))
	for x := 0; x < 16; x++ {
		for y := 0; y < 8; y++ {
			img.Set(x, y, color.White)
		}
	}
	var src image.Image = img
	disableDithering = true
	defer func() { disableDithering = false }()

	result := ImgToBytes(16, 8, &src)
	if len(result) != 16 {
		t.Fatalf("expected 16 bytes, got %d", len(result))
	}
	// All white pixels should produce all-zero bits
	for i, b := range result {
		if b != 0x00 {
			t.Errorf("byte %d: expected 0x00, got 0x%02X", i, b)
		}
	}
}

func TestEncodeToString(t *testing.T) {
	input := []byte{0x00, 0xFF, 0x42}
	result := EncodeToString(input)
	if result == "" {
		t.Error("expected non-empty base64 string")
	}
	// Known base64 encoding
	expected := "AP9C"
	if result != expected {
		t.Errorf("EncodeToString(%v) = %q, want %q", input, result, expected)
	}
}

func TestWriteToBinFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.bin")
	data := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	if err := WriteToBinFile(path, data); err != nil {
		t.Fatalf("WriteToBinFile: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if len(got) != len(data) {
		t.Fatalf("expected %d bytes, got %d", len(data), len(got))
	}
	for i := range data {
		if got[i] != data[i] {
			t.Errorf("byte %d: expected 0x%02X, got 0x%02X", i, data[i], got[i])
		}
	}
}

func TestWriteToGoFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test-generated.go")
	data := []byte{0x01, 0x02}
	if err := WriteToGoFile(path, "testvar", data); err != nil {
		t.Fatalf("WriteToGoFile: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	content := string(got)
	if len(content) == 0 {
		t.Fatal("generated file is empty")
	}
	// Should contain the variable name and package declaration
	if !contains(content, "package main") {
		t.Error("missing package declaration")
	}
	if !contains(content, "rtestvar") {
		t.Error("missing variable name")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestLoadImg_NotFound(t *testing.T) {
	_, err := LoadImg("/nonexistent/path/to/image.png")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}
