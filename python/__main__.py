import pulumi
from pulumi_azure_native import resources
import pierskarsenbarg_storage as storage

resource_group = resources.ResourceGroup("resource_group")

my_storage = storage.Storage("storage",
                             resource_group_name=resource_group.name
                             )

pulumi.export("blob_container_name", my_storage.blob_container_name)
