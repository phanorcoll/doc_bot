// Package main provides main  
package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// PluginInfo struct  
type PluginInfo struct {
	Title string
	Desc  string
	URL   string
}

func main() {
	// Define flags with default values and help messages
	pluginDir := flag.String("dir", "", "Directory containing Neovim plugin files (required)")
	outputPath := flag.String("output", "README.md", "Output file path for the generated documentation")
	templatePath := flag.String("template", "template.tmpl", "Path to the template file for formatting")
	flag.Parse()

	if *pluginDir == "" {
		fmt.Println("Error: Missing required flag --dir")
		flag.PrintDefaults()
		return
	}

	var plugins []PluginInfo

	err := filepath.Walk(*pluginDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info, ok := info.(os.FileInfo); ok && !info.IsDir() {
			fmt.Printf("Parsing file: %s\n", path)
			pluginInfo, err := parsePluginFile(path)
			if err != nil {
				fmt.Printf("Error parsing file %s: %v\n", path, err)
				return nil
			}

			if pluginInfo != nil {
				plugins = append(plugins, *pluginInfo)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking directory:", err)
		return
	}

	if len(plugins) == 0 {
		fmt.Println("No valid plugin files found in", *pluginDir)
		return
	}

	generateReadme(plugins, *outputPath, *templatePath)
	fmt.Println("README file generated successfully!")
}

func parsePluginFile(path string) (*PluginInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Define a regular expression to match header lines
	re := regexp.MustCompile(`^\s*--(?:\s*url: (.*?))?(?:\s*desc: (.*?))?\s*$`)

	var info PluginInfo

	for _, line := range strings.Split(string(data), "\n") {
		match := re.FindStringSubmatch(line)
		if match != nil {
			if len(match) > 1 && match[1] != "" {
				info.URL = strings.TrimSpace(match[1])
				// Extract package name from url
				parts := strings.Split(info.URL, "/")
				if len(parts) < 3 {
					return nil, fmt.Errorf("invalid url format in %s", path)
				}
				info.Title = parts[len(parts)-2] + "/" + parts[len(parts)-1]
			}
			if len(match) > 2 && match[2] != "" {
				info.Desc = strings.TrimSpace(match[2])
			}
		}
	}

	// Only return a non-nil PluginInfo if it has a URL
	if info.URL == "" {
		return nil, nil
	}

	return &info, nil
}

func generateReadme(plugins []PluginInfo, outputPath, templatePath string) {
	// Load the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Create or open the README file in append mode
	f, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Get file info to check size
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	fmt.Println("Error getting file info:", err)
	// 	return
	// }
	//
	// // If the file is empty, add the header
	// if fileInfo.Size() == 0 {
	// 	if err != nil {
	// 		fmt.Println("Error writing to file:", err)
	// 		return
	// 	}
	// }

	// Execute the template and write to the file
	err = tmpl.Execute(f, plugins)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
