# Azure Storage component examples

This repo holds examples of Pulumi programs that reference the [Storage component](https://github.com/pierskarsenbarg/storage).

There are examples for the following languages:

* [YAML](./yaml/README.md)
* [Python](./python/README.md)
* [Typescript](./typescript/README.md)

(The above links will take you through to the README for that language which have instructions on how to run each example program)

> [!NOTE]
> Since the component is written in Python, there is a requirement to install Python on the machine where you're running this code. If you don't have Python installed, you can follow the instructions in the [Pulumi Docs](https://www.pulumi.com/docs/iac/languages-sdks/python/). 

## Pre-Requisites for all examples

All of the examples have the following pre-requistes. Any language specific pre-requisites will be shown in the README text for each of the links above.

* [Pulumi CLI (>3.159.0)](https://www.pulumi.com/docs/iac/download-install/) installed
* [Python3](https://www.pulumi.com/docs/iac/languages-sdks/python/) installed
* [State Backend](https://www.pulumi.com/docs/iac/concepts/state-and-backends/) - I recommend [Pulumi Cloud](https://app.pulumi.com) but this will work across all options of state backends
* [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli) - you should also login to a valid subscription
