package main

import (
	"context"
	"log"
	"os"

	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/apache/arrow/go/v17/parquet"
	"github.com/apache/arrow/go/v17/parquet/compress"
	"github.com/apache/arrow/go/v17/parquet/file"
	"github.com/apache/arrow/go/v17/parquet/pqarrow"
)

func main() {
	// Inicializa o alocador de memória padrão
	alloc := memory.DefaultAllocator

	// Abre o arquivo Parquet de entrada
	inputFile, err := os.Open("sales.parquet")
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de entrada: %v", err)
	}
	defer inputFile.Close()

	// Cria um leitor de Parquet
	pf, err := file.NewParquetReader(inputFile)
	if err != nil {
		log.Fatalf("Erro ao criar o leitor de Parquet: %v", err)
	}
	defer pf.Close()

	// Cria um leitor de arquivo pqarrow
	reader, err := pqarrow.NewFileReader(pf, pqarrow.ArrowReadProperties{}, alloc)
	if err != nil {
		log.Fatalf("Erro ao criar o leitor pqarrow: %v", err)
	}

	// Lê o conteúdo do arquivo Parquet em uma tabela Arrow
	ctx := context.Background()
	table, err := reader.ReadTable(ctx)
	if err != nil {
		log.Fatalf("Erro ao ler a tabela: %v", err)
	}
	defer table.Release()

	// Abre o arquivo Parquet de saída
	outputFile, err := os.Create("output.parquet")
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo de saída: %v", err)
	}
	defer outputFile.Close()

	// Define as propriedades do escritor
	writerProps := parquet.NewWriterProperties(parquet.WithCompression(compress.Codecs.Snappy))
	arrProps := pqarrow.DefaultWriterProps()

	// Escreve a tabela Arrow no novo arquivo Parquet
	if err := pqarrow.WriteTable(table, outputFile, table.NumRows(), writerProps, arrProps); err != nil {
		log.Fatalf("Erro ao escrever a tabela: %v", err)
	}

	log.Println("Arquivo Parquet processado com sucesso.")
}
