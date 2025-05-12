package main

import (
	storage "github.com/pierskarsenbarg/storage-component-examples/go/sdks/pierskarsenbarg-storage/storage"

	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}

		_, err = storage.NewStorage(ctx, "mystorage", &storage.StorageArgs{
			ResourceGroupName: resourceGroup.Name,
		})

		return nil
	})
}
