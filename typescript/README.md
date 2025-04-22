# Azure Storage Component Typescript example

This is an example Pulumi program in Typescript that will use the Storage component to create the necessary resources.

## Pre-requisites

* [NodeJS](https://www.pulumi.com/docs/iac/languages-sdks/javascript/)

## Instructions

1. Clone this repo: `git clone https://github.com/pierskarsenbarg/storage-component-examples`
1. Change directory: `cd typescript`
1. Install dependencies: `pulumi install` (this will also run `npm install`)
1. Set location to create the resources in: `pulumi config set azure-native:location {location}`
1. Run update: `pulumi up`

## Destroy

Don't forget to do this so that you don't get hit with any costs.

1. `pulumi destroy`
