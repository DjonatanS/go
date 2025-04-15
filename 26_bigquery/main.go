package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "data-project-454013")
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}
	defer client.Close() // Good practice to close the client

	// Example usage of the BigqueryStruct

	bq := Dataframe{
		Read: &BigqueryStruct{
			ctx: ctx,
		},
	}

	projectID := "data-project-454013"
	datasetID := "teste"
	tableID := "colaboradores"
	headers, err := bq.Read.getHeaders(projectID, datasetID, tableID)
	if err != nil {
		fmt.Println("Error getting headers:", err)
		return
	}
	fmt.Println("Headers:", headers)

}

// Dataframe represents a data structure that can read data from BigQuery.
type Dataframe struct {
	Read read
}
type BigqueryStruct struct {
	ctx context.Context
}

type read interface {
	getData(projectID, datasetID, tableID string) ([][]string, error)
	getHeaders(projectID, datasetID, tableID string) ([]string, error)
}

func (b *BigqueryStruct) getData(projectID, datasetID, tableID string) ([][]string, error) {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar cliente BigQuery: %w", err)
	}
	defer client.Close()

	q := client.Query(fmt.Sprintf("SELECT * FROM `%s.%s.%s`", projectID, datasetID, tableID))
	it, err := q.Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("falha ao executar a query: %w", err)
	}

	var data [][]string
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("erro ao iterar sobre os resultados: %w", err)
		}

		var row []string
		for _, v := range values {
			row = append(row, fmt.Sprintf("%v", v))
		}
		data = append(data, row)
	}

	return data, nil
}

func (b *BigqueryStruct) getHeaders(projectID, datasetID, tableID string) ([]string, error) {
	ctx := context.Background()

	// 1. Cria o cliente BigQuery
	// O cliente usará as Application Default Credentials (ADC).
	// Certifique-se de que você está autenticado (ex: gcloud auth application-default login)
	// ou que as variáveis de ambiente apropriadas (GOOGLE_APPLICATION_CREDENTIALS) estão definidas.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar cliente BigQuery: %w", err)
	}
	defer client.Close() // É importante fechar o cliente quando terminar

	// 2. Obtém a referência da tabela
	tableRef := client.Dataset(datasetID).Table(tableID)

	// 3. Busca os metadados da tabela
	// A informação do esquema (incluindo nomes das colunas) está nos metadados.
	meta, err := tableRef.Metadata(ctx)
	if err != nil {
		// Você pode querer tratar erros específicos aqui, como bigquery.ErrNotFound
		return nil, fmt.Errorf("falha ao obter metadados da tabela '%s': %w", tableID, err)
	}

	// 4. Extrai os nomes das colunas do esquema
	var headers []string
	if meta.Schema == nil {
		// Caso a tabela não tenha um esquema definido (raro para tabelas padrão)
		return headers, fmt.Errorf("a tabela '%s' não parece ter um esquema definido", tableID)
	}

	// O Schema é um slice de *bigquery.FieldSchema
	for _, fieldSchema := range meta.Schema {
		headers = append(headers, fieldSchema.Name) // O campo 'Name' contém o nome da coluna
	}

	return headers, nil
}

// Query executes a BigQuery query and returns headers, rows, and an error.
func Query(ctx context.Context, client *bigquery.Client, query string) ([]bigquery.Value, [][]bigquery.Value, error) {
	q := client.Query(query)
	it, err := q.Read(ctx)
	if err != nil {
		return nil, nil, err
	}
	// schema := it.Schema // schema is used below directly
	// fmt.Println("schema:", schema)
	var rows [][]bigquery.Value
	var headers []bigquery.Value
	for _, field := range it.Schema {
		// fmt.Println("field:", field)
		// fmt.Println("field.Name:", field.Name)
		headers = append(headers, field.Name)
	}
	// fmt.Println("headers:", headers) // Headers are now returned
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		rows = append(rows, values)
	}

	return headers, rows, nil
}
