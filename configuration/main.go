package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type ConfigurationItem struct {
	Name           string
	Rpc            string
	ChainId        string
	Symbol         string
	IpfsGateway    string
	LocalExplorer  string
	RemoteExplorer string
}

type ConfigurationPost []ConfigurationItem

func EnvsFromConfiguration(item ConfigurationItem) string {
	var b strings.Builder

	prefix := fmt.Sprintf("TB_CHAINS_%s_", strings.ToUpper(item.Name))
	b.WriteString(prefix + "CHAINID=" + item.ChainId + "\n")
	b.WriteString(prefix + "RPCPROVIDER=" + item.Rpc + "\n")
	b.WriteString(prefix + "SYMBOL=" + item.Symbol + "\n")
	b.WriteString(prefix + "PINGATEWAY=" + item.IpfsGateway + "\n")
	b.WriteString(prefix + "LOCALEXPLORER=" + item.LocalExplorer + "\n")
	b.WriteString(prefix + "REMOTEEXPLORER=" + item.RemoteExplorer + "\n")

	return b.String()
}

func WriteEnvFile(path string, contents string) (err error) {
	file, err := os.Create(path)
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return
	}

	_, err = file.WriteString(contents)
	if err != nil {
		return
	}

	return nil
}

func SaveConfiguration(path string, items []ConfigurationItem) (err error) {
	lines := []string{}
	for _, item := range items {
		lines = append(lines, EnvsFromConfiguration(item))
	}
	err = WriteEnvFile(path, strings.Join(lines, "\n"))
	return err
}

func SaveJson(path string, items []ConfigurationItem) (err error) {
	contents, err := json.MarshalIndent(items, "", "	")
	if err != nil {
		return
	}

	file, err := os.Create(path)
	if err != nil {
		return
	}

	file.Write(contents)
	return nil
}

func ReadJson(path string) (contents []byte, err error) {
	contents, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	return
}

func makeConfigurationHandler(outputDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodGet {
			http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
			return
		}

		if r.Method == http.MethodPost {
			log.Println("POST /configuration")

			p := make([]ConfigurationItem, 0)

			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				log.Println("error", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = SaveConfiguration(path.Join(outputDir, "configuration.env"), p)
			if err != nil {
				http.Error(w, fmt.Sprintf("Could not save configuration file: %s", err), http.StatusInternalServerError)
				return
			}
			err = SaveJson(path.Join(outputDir, "configuration.json"), p)
			if err != nil {
				http.Error(w, fmt.Sprintf("Could not save JSON configuration file: %s", err), http.StatusInternalServerError)
				return
			}
		}

		if r.Method == http.MethodGet {
			log.Println("GET /configuration")

			data, err := ReadJson(path.Join(outputDir, "configuration.json"))
			if err != nil {
				http.Error(w, fmt.Sprintf("Could not read JSON configuration file: %s", err), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	}
}

func main() {
	var outputPath string

	flag.StringVar(&outputPath, "dir", "/output/", "Path to output directory")
	flag.Parse()
	log.Println("Will save output to", outputPath)

	http.HandleFunc("/configuration", makeConfigurationHandler(outputPath))

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
