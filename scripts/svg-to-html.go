package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run svg-to-html.go <input-file.ts> <output-file.html>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Read error:", err)
		os.Exit(1)
	}

	svgPaths, err := extractSVGPaths(content)
	if err != nil {
		fmt.Println("Extraction error:", err)
		os.Exit(1)
	}

	svgContent := buildSVGContent(svgPaths)

	if err := writeFile(outputFile, svgContent); err != nil {
		fmt.Println("Write error:", err)
		os.Exit(1)
	}

	fmt.Println("✅ SVG HTML created successfully:", outputFile)
}

// readFile reads the content of the input file.
func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// extractSVGPaths extracts SVG path data from CoreUI TypeScript file.
func extractSVGPaths(content string) ([]string, error) {
	// Step 1: Capture the SVG string within double quotes after the size definition
	re := regexp.MustCompile(`(?s)"(?:\d+\s\d+)?","(.*?)"`)
	mainMatch := re.FindStringSubmatch(content)
	if len(mainMatch) < 2 {
		return nil, fmt.Errorf("no SVG data found in input file")
	}

	svgContent := mainMatch[1]

	// Step 2: Extract individual path elements from the captured SVG content
	pathRe := regexp.MustCompile(`<path\s+d=['"]([^'"]+)['"].*?/?\s*>`)
	pathMatches := pathRe.FindAllStringSubmatch(svgContent, -1)

	if len(pathMatches) == 0 {
		return nil, fmt.Errorf("no SVG paths found in extracted SVG content")
	}

	var paths []string
	for _, pmatch := range pathMatches {
		paths = append(paths, pmatch[1])
	}

	return paths, nil
}

// buildSVGContent constructs the SVG HTML content from path data.
func buildSVGContent(paths []string) string {
	svg := "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 32 32\">\n"
	for _, path := range paths {
		svg += fmt.Sprintf("  <path d='%s'/>\n", path)
	}
	svg += "</svg>\n"
	return svg
}

// writeFile writes the SVG content to the specified output file.
func writeFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
