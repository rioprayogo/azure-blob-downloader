package azure

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func DownloadFolderFromAzure(containerName, folderPrefix, connectionString, outputDir string) error {
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		return fmt.Errorf("failed to create azblob client: %v", err)
	}

	pager := client.NewListBlobsFlatPager(containerName, &azblob.ListBlobsFlatOptions{
		Prefix: &folderPrefix,
	})

	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			return fmt.Errorf("failed to list blobs: %v", err)
		}

		for _, blob := range page.Segment.BlobItems {
			blobName := *blob.Name
			fmt.Printf("Downloading blob: %s\n", blobName)

			downloadResponse, err := client.DownloadStream(context.TODO(), containerName, blobName, nil)
			if err != nil {
				return fmt.Errorf("failed to download blob %s: %v", blobName, err)
			}
			defer downloadResponse.Body.Close()

			relativePath := strings.TrimPrefix(blobName, folderPrefix)
			outputPath := filepath.Join(outputDir, relativePath)

			if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
				return fmt.Errorf("failed to create directories for %s: %v", outputPath, err)
			}

			outputFile, err := os.Create(outputPath)
			if err != nil {
				return fmt.Errorf("failed to create output file %s: %v", outputPath, err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, downloadResponse.Body)
			if err != nil {
				return fmt.Errorf("failed to save blob %s to file: %v", blobName, err)
			}

			fmt.Printf("Blob %s downloaded to %s\n", blobName, outputPath)
		}
	}

	return nil
}
