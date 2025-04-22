# Azure Storage Component Python example

This is an example Pulumi program in Python that will use the Storage component to create the necessary resources.

## Pre-requisites

There are no language specific pre-requisites.

## Instructions

1. Clone this repo: `git clone https://github.com/pierskarsenbarg/storage-component-examples`
1. Change directory: `cd python`
1. Set up your virtual environment: `python3 -m venv venv`
1. Activate the virtual environment: `source venv/bin/activate`
1. Install dependencies: `pulumi install` (this will also install the packages )
1. Set location to create the resources in: `pulumi config set azure-native:location {location}`
1. Run update: `pulumi up`

## Destroy

Don't forget to do this so that you don't get hit with any costs.

1. `pulumi destroy`
