package services

import (
	"fmt"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const webpExt = ".webp"

func DirWalk(path string) {
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if isWebpFile(path) {
				log.Println("isWebpFile")
				r, err := readFile(path)
				if err != nil {
					return err
				}

				rp, err := decodeWebp(r)
				if err != nil {
					return err
				}

				nf, err := buildJpegFile(buildJpegFileName(path))
				if err != nil {
					return err
				}

				return writeJpeg(nf, rp)
			}

			return nil
		})
	
	if err != nil {
		log.Println(err)
	}
}

func isWebpFile(path string) bool {
	ext := filepath.Ext(path)
	if ext == webpExt {
		return true
	}

	return false
}

func readFile(path string) (io.Reader, error) {
	log.Println("readFile")
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func decodeWebp(r io.Reader) (image.Image, error) {
	log.Println("decodeWebp")
	return webp.Decode(r)
}

func buildJpegFileName(f string) string {
	log.Println("buildJpegFileName")
	return fmt.Sprintf("%v.jpeg", strings.Split(f, ".")[0])
}

func buildJpegFile(p string) (io.Writer, error) {
	log.Println("buildJpegFile")
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func writeJpeg(w io.Writer, i image.Image) error {
	log.Println("writeJpeg")
	var rgba *image.RGBA
	if nrgba, ok := i.(*image.NRGBA); ok {
		if nrgba.Opaque() {
			rgba = &image.RGBA{
				Pix:    nrgba.Pix,
				Stride: nrgba.Stride,
				Rect:   nrgba.Rect,
			}
		}
	}

	if rgba != nil {
		return jpeg.Encode(w, rgba, &jpeg.Options{Quality: 95})
	} else {
		return jpeg.Encode(w, i, &jpeg.Options{Quality: 100})
	}
}
