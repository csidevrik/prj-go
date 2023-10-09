package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	folderPath := "C:\\Users\\adminos\\OneDrive\\2A-JOB02-EMOVEP\\2023\\CONTRATOS\\RE-EP-EMOVEP-2023-02\\FACTURAS\\SEP\\RDD" // Cambia esto a la ruta de tu carpeta

	// Crear un mapa para rastrear los hashes de los archivos
	fileHashes := make(map[string][]string)

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileHash, err := hashFile(path)
			if err != nil {
				return err
			}
			fileHashes[fileHash] = append(fileHashes[fileHash], path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error al recorrer la carpeta:", err)
		return
	}

	// Eliminar archivos duplicados
	for _, paths := range fileHashes {
		if len(paths) > 1 {
			fmt.Printf("Archivos duplicados:/n")
			for i := 1; i < len(paths); i++ {
				fmt.Println(paths[i])
				err := os.Remove(paths[i])
				if err != nil {
					fmt.Println("Error al eliminar el archivo:", err)
				}
			}
		}
	}
}

func hashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}
