package main

import (
	"fmt"
	"golang-azure-download/internal/azure"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: <container_name> <connection_string> <output_path> <folder_prefix1> [<folder_prefix2> ...]")
		os.Exit(1)
	}

	containerName := os.Args[1]
	connectionString := os.Args[2]
	outputPath := os.Args[3]
	folderPrefixes := os.Args[4:]

	for _, folderPrefix := range folderPrefixes {
		fmt.Printf("Downloading folder: %s\n", folderPrefix)
		err := azure.DownloadFolderFromAzure(containerName, folderPrefix, connectionString, outputPath)
		if err != nil {
			fmt.Printf("Error downloading folder %s: %v\n", folderPrefix, err)
			os.Exit(1)
		}
	}

	fmt.Println("All folders downloaded successfully.")
}
