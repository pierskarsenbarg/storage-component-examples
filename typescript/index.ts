import * as pulumi from "@pulumi/pulumi";
import * as resources from "@pulumi/azure-native/resources";
import * as storage from "@pierskarsenbarg/storage";

const resourceGroup = new resources.ResourceGroup("pk-rg")

const myStorage = new storage.Storage("mystorage", {
    resourceGroupName: resourceGroup.name
})

export const blobContainerName = myStorage.blobContainerName;