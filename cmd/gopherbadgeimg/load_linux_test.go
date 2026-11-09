package main

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"runtime/debug"
	"testing"
)

func TestLoadImgClosesFile(t *testing.T) {
	// A finalizer must not hide a missing Close in LoadImg.
	previousGC := debug.SetGCPercent(-1)
	defer debug.SetGCPercent(previousGC)

	for _, valid := range []bool{true, false} {
		name := "invalid"
		if valid {
			name = "valid"
		}
		t.Run(name, func(t *testing.T) {
			filename := filepath.Join(t.TempDir(), "image.png")
			file, err := os.Create(filename)
			if err != nil {
				t.Fatal(err)
			}
			if valid {
				err = png.Encode(file, image.NewRGBA(image.Rect(0, 0, 1, 1)))
			} else {
				_, err = file.WriteString("not an image")
			}
			closeErr := file.Close()
			if err != nil {
				t.Fatal(err)
			}
			if closeErr != nil {
				t.Fatal(closeErr)
			}
			img, err := LoadImg(filename)
			if valid {
				if err != nil {
					t.Fatal(err)
				}
				if (*img).Bounds() != image.Rect(0, 0, 1, 1) {
					t.Fatal("unexpected image bounds")
				}
			} else if err == nil {
				t.Fatal("expected decode failure")
			}
			descriptors, err := os.ReadDir("/proc/self/fd")
			if err != nil {
				t.Fatal(err)
			}
			for _, descriptor := range descriptors {
				target, err := os.Readlink(filepath.Join("/proc/self/fd", descriptor.Name()))
				if err == nil && target == filename {
					t.Errorf("LoadImg left %s open", filename)
				}
			}
		})
	}
}
