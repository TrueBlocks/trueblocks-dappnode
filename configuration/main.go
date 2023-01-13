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
	"strconv"
	"strings"
)

type ConfigurationItem struct {
	Name           string
	Rpc            string
	ChainId        string
	Symbol         string
	// IpfsGateway    string
	LocalExplorer  string
	RemoteExplorer string
	ScraperArgs    string
	ScraperFile    string
}

type GlobalConfiguration struct {
	RunScraper   bool
	InitBlooms   bool
	InitIndex    bool
	MonitorArgs  string
	MonitorFile  string
	EtherscanKey string
	DefaultGateway string
}

type ConfigurationPost struct {
	Global GlobalConfiguration
	Chains []ConfigurationItem
}

func EnvsFromConfiguration(item ConfigurationItem) string {
	var b strings.Builder

	prefix := fmt.Sprintf("export TB_CHAINS_%s_", strings.ToUpper(item.Name))
	b.WriteString(prefix + "CHAINID=" + item.ChainId + "\n")
	b.WriteString(prefix + "RPCPROVIDER=" + item.Rpc + "\n")
	b.WriteString(prefix + "SYMBOL=" + item.Symbol + "\n")

	// if item.IpfsGateway != "" {
	// 	b.WriteString(prefix + "PINGATEWAY=" + item.IpfsGateway + "\n")
	// }
	if item.LocalExplorer != "" {
		b.WriteString(prefix + "LOCALEXPLORER=" + item.LocalExplorer + "\n")
	}
	if item.RemoteExplorer != "" {
		b.WriteString(prefix + "REMOTEEXPLORER=" + item.RemoteExplorer + "\n")
	}
	if item.ScraperArgs != "" {
		b.WriteString(fmt.Sprintf("export SCRAPER_%s_ARGS=%s\n", strings.ToUpper(item.Name), normalizeUserInput(item.ScraperArgs)))
	}
	if item.ScraperFile != "" {
		b.WriteString(fmt.Sprintf("export SCRAPER_%s_FILE=%s\n", strings.ToUpper(item.Name), normalizeUserInput(item.ScraperFile)))
	}

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

func normalizeUserInput(content string) string {
	return strconv.Quote(content)
}

func SaveConfiguration(path string, config ConfigurationPost) (err error) {
	lines := []string{
		"#!/bin/bash",
		fmt.Sprint("export RUN_SCRAPER=", config.Global.RunScraper),
		fmt.Sprint("export BOOTSTRAP_BLOOM_FILTERS=", config.Global.InitBlooms),
		fmt.Sprint("export BOOTSTRAP_FULL_INDEX=", config.Global.InitIndex),
		fmt.Sprint("export TB_KEYS_ETHERSCAN_APIKEY=", normalizeUserInput(config.Global.EtherscanKey)),
		fmt.Sprint("export TB_SETTINGS_DEFAULTGATEWAY=", config.Global.DefaultGateway),
	}
	if config.Global.MonitorArgs != "" {
		lines = append(lines, fmt.Sprint("export MONITORS_ARGS=", normalizeUserInput(config.Global.MonitorArgs)))
	}
	if config.Global.MonitorFile != "" {
		lines = append(lines, fmt.Sprint("export MONITORS_FILE=", normalizeUserInput(config.Global.MonitorFile)))
	}

	for _, item := range config.Chains {
		lines = append(lines, EnvsFromConfiguration(item))
	}
	err = WriteEnvFile(path, strings.Join(lines, "\n"))
	return err
}

func SaveJson(path string, config ConfigurationPost) (err error) {
	contents, err := json.MarshalIndent(config, "", "	")
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

			p := ConfigurationPost{}

			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				log.Println("error", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = SaveConfiguration(path.Join(outputDir, "configuration.sh"), p)
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
	var port string
	var staticDir string

	flag.StringVar(&outputPath, "dir", "/output/", "Path to output directory")
	flag.StringVar(&port, "port", "80", "Port to listen on")
	flag.StringVar(&staticDir, "static", "./static", "Directory to serve static files from")
	flag.Parse()
	log.Println("Will save output to", outputPath)

	http.HandleFunc("/configuration", makeConfigurationHandler(outputPath))

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)

	log.Print("Listening on ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
