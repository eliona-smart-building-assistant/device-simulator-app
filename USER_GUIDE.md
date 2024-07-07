<mark>
Replace the following strings globally with the real Device Simulator:
</mark>

- `device-simulator` with app's name. Use `-` if the Device Simulator contains spaces. All letters are lower-cased.
- `device_simulator` and `device\\_simulator` with app's name. Use `_` if the Device Simulator contains spaces. All letters are lower-cased.
- `Device Simulator` with app's real names. Letters can be mixed-cased, depends on using e.g. brand names.

# Device Simulator User Guide

### Introduction

> The Device Simulator app provides integration and synchronization between Eliona and Device Simulator services.

## Overview

This guide provides instructions on configuring, installing, and using the Device Simulator app to manage resources and synchronize data between Eliona and Device Simulator services.

## Installation

Install the Device Simulator app via the Eliona App Store.

## Configuration

The Device Simulator app requires configuration through Eliona’s settings interface. Below are the general steps and details needed to configure the app effectively.

### Registering the app in Device Simulator Service

Create credentials in Device Simulator Service to connect the Device Simulator services from Eliona. All required credentials are listed below in the [configuration section](#configure-the-device-simulator-app).  

<mark>TODO: Describe the steps where you can get or create the necessary credentials.</mark> 

### Configure the Device Simulator app 

Configurations can be created in Eliona under `Apps > Device Simulator > Settings` which opens the app's [Generic Frontend](https://doc.eliona.io/collection/v/eliona-english/manuals/settings/apps). Here you can use the appropriate endpoint with the POST method. Each configuration requires the following data:

| Attribute         | Description                                                                     |
|-------------------|---------------------------------------------------------------------------------|
| `baseURL`         | URL of the Device Simulator services.                                                   |
| `clientSecrets`   | Client secrets obtained from the Device Simulator service.                              |
| `assetFilter`     | Filtering asset during [Continuous Asset Creation](#continuous-asset-creation). |
| `enable`          | Flag to enable or disable this configuration.                                   |
| `refreshInterval` | Interval in seconds for data synchronization.                                   |
| `requestTimeout`  | API query timeout in seconds.                                                   |
| `projectIDs`      | List of Eliona project IDs for data collection.                                 |

Example configuration JSON:

```json
{
  "baseURL": "http://service/v1",
  "clientSecrets": "random-cl13nt-s3cr3t",
  "filter": "",
  "enable": true,
  "refreshInterval": 60,
  "requestTimeout": 120,
  "projectIDs": [
    "10"
  ]
}
```

## Continuous Asset Creation

Once configured, the app starts Continuous Asset Creation (CAC). Discovered resources are automatically created as assets in Eliona, and users are notified via Eliona’s notification system.

<mark>TODO: Describe what resources are created, the hierarchy and the data points.</mark>

## Additional Features

<mark>TODO: Describe all other features of the app.</mark>

### Dashboard templates

The app offers a predefined dashboard that clearly displays the most important information. YOu can create such a dashboard under `Dashboards > Copy Dashboard > From App > Device Simulator`.

### <mark>TODO: Other features</mark>
