package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/disintegration/imaging"
)

var verbose bool

// Función para imprimir sólo cuando la opción verbose esté activa
func logVerbose(args ...interface{}) {
	if verbose {
		fmt.Println(args...)
	}
}

func resizeImage(inputPath, outputFolder string, width, height int) {
	// Abrir una imagen.
	src, err := imaging.Open(inputPath)
	if err != nil {
		fmt.Printf("failed to open image: %v\n", err)
		return
	}

	// Redimensionar la imagen.
	dstImage := imaging.Fill(src, width, height, imaging.Center, imaging.Lanczos)

	// Construir el nombre del archivo de salida.
	filename := filepath.Base(inputPath)
	nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
	newFilename := fmt.Sprintf("red_%d_%d_%s.jpg", width, height, nameWithoutExt)
	outputPath := filepath.Join(outputFolder, newFilename)

	// Guardar la imagen redimensionada.
	err = imaging.Save(dstImage, outputPath, imaging.JPEGQuality(95))
	if err != nil {
		logVerbose("Failed to save image:", err)
	} else {
		logVerbose("Image saved successfully:", outputPath)
	}
}

func main() {
	// Define las banderas (flags)
	widthPtr := flag.Int("w", 100, "Width of the resized image")
	widthPtrLong := flag.Int("width", 100, "Width of the resized image (long form)")
	heightPtr := flag.Int("h", 100, "Height of the resized image")
	heightPtrLong := flag.Int("height", 100, "Height of the resized image (long form)")
	filePtr := flag.String("f", "", "Specific file to resize")
	filePtrLong := flag.String("file", "", "Specific file to resize (long form)")
	pathPtr := flag.String("p", "./imgs", "Directory path to resize all images in it")
	pathPtrLong := flag.String("path", "./imgs", "Directory path to resize all images in it (long form)")
	flag.BoolVar(&verbose, "v", false, "Enable verbose mode")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose mode (long form)")

	// Parsear los argumentos
	flag.Parse()

	// Determinar el valor final de los parámetros
	width := *widthPtr
	if *widthPtrLong != 100 {
		width = *widthPtrLong
	}

	height := *heightPtr
	if *heightPtrLong != 100 {
		height = *heightPtrLong
	}

	file := *filePtr
	if *filePtrLong != "" {
		file = *filePtrLong
	}

	path := *pathPtr
	if *pathPtrLong != "./imgs" {
		path = *pathPtrLong
	}

	outputFolder := path

	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		os.MkdirAll(outputFolder, os.ModePerm)
		logVerbose("Created output folder:", outputFolder)
	}

	re := regexp.MustCompile(`^red_\d+_\d+_.*`)

	// Si se especifica un archivo, solo procesa ese archivo
	if file != "" {
		outputFolder = filepath.Dir(file)
		if !re.MatchString(file) {
			resizeImage(file, outputFolder, width, height)
			logVerbose("Resizing single file:", file)
		} else {
			logVerbose("File already resized, skipping:", file)
		}
		return
	}

	// Procesa todos los archivos en el directorio especificado
	files, _ := os.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if re.MatchString(file.Name()) {
			continue
		}

		lowerName := strings.ToLower(file.Name())
		if strings.HasSuffix(lowerName, ".png") || strings.HasSuffix(lowerName, ".jpg") ||
			strings.HasSuffix(lowerName, ".jpeg") || strings.HasSuffix(lowerName, ".bmp") ||
			strings.HasSuffix(lowerName, ".tiff") || strings.HasSuffix(lowerName, ".webp") {
			inputPath := filepath.Join(path, file.Name())
			// fmt.Println(file.Name(), inputPath)
			resizeImage(inputPath, outputFolder, width, height)
		}
	}
}
