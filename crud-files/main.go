package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FileOperationRequest struct {
	Operation string          `json:"operation"`
	Data      interface{}     `json:"data,omitempty"`
	Filename  string          `json:"filename"`
	Format    string          `json:"format"`
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// ファイル操作を行う関数
func processFile(req FileOperationRequest) (interface{}, error) {
	switch req.Operation {
	case "read":
		return readFile(req.Filename, req.Format)
	case "write":
		return nil, writeFile(req.Filename, req.Format, req.Data)
	case "append":
		return nil, appendToFile(req.Filename, req.Format, req.Data)
	case "list":
		return listFiles(req.Filename)
	default:
		return nil, nil
	}
}

func readFile(filename, format string) (interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	switch format {
	case "json":
		var data interface{}
		if err := json.NewDecoder(file).Decode(&data); err != nil {
			return nil, err
		}
		return data, nil
	case "csv":
		reader := csv.NewReader(file)
		var records [][]string
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			records = append(records, record)
		}
		return records, nil
	default:
		content, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		return string(content), nil
	}
}

func writeFile(filename, format string, data interface{}) error {
	// ディレクトリが存在しない場合は作成
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "json":
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		return encoder.Encode(data)
	case "csv":
		writer := csv.NewWriter(file)
		defer writer.Flush()

		switch v := data.(type) {
		case [][]string:
			return writer.WriteAll(v)
		case []Person:
			// ヘッダーを書き込み
			if err := writer.Write([]string{"Name", "Age", "Email"}); err != nil {
				return err
			}
			// データを書き込み
			for _, person := range v {
				record := []string{
					person.Name,
					strconv.Itoa(person.Age),
					person.Email,
				}
				if err := writer.Write(record); err != nil {
					return err
				}
			}
		}
		return nil
	default:
		_, err := file.WriteString(data.(string))
		return err
	}
}

func appendToFile(filename, format string, data interface{}) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "json":
		// JSONファイルの場合は、既存のデータを読み込んで結合する必要がある
		var existingData []interface{}
		if fileInfo, err := file.Stat(); err == nil && fileInfo.Size() > 0 {
			if err := json.NewDecoder(file).Decode(&existingData); err != nil {
				return err
			}
		}
		existingData = append(existingData, data)
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		return encoder.Encode(existingData)
	case "csv":
		writer := csv.NewWriter(file)
		defer writer.Flush()

		switch v := data.(type) {
		case []string:
			return writer.Write(v)
		case Person:
			record := []string{v.Name, strconv.Itoa(v.Age), v.Email}
			return writer.Write(record)
		}
		return nil
	default:
		_, err := file.WriteString(data.(string))
		return err
	}
}

func listFiles(dir string) (interface{}, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req FileOperationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := processFile(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Message: "File operation completed",
		Data:    result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/file", fileHandler)

	log.Println("File operations server starting on :8083...")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
} 