package main

import (
	"go-batch/processor"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	log.Println("Iniciando el procesamiento de datos")
	//	db.InitDB()

	totalRecords := 100000 // Número de registros
	recordSize := 1024     // Tamaño en bytes (1 KB por registro)
	batchSize := 500       // Tamaño de lotes a insertar
	workerCount := 10      // Número de workers concurrentes

	log.Printf("Generando %d registros de %d bytes cada uno", totalRecords, recordSize)
	data := GenerateData(totalRecords, recordSize)

	log.Println("Iniciando procesamiento...")
	start := time.Now()

	// Ejecutar el procesamiento
	processor.WorkerPool(data, batchSize, workerCount)

	elapsed := time.Since(start)
	log.Printf("Procesamiento finalizado en %s: ", elapsed)
}

func GenerateData(recordCount int, recordSize int) []string {
	data := make([]string, recordCount)
	for i := 0; i < recordCount; i++ {
		data[i] = generateRandomString(recordSize)
	}
	return data
}

func generateRandomString(size int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for i := 0; i < size; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
