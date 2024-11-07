# Device Simulator User Guide

### Introduction

> The Device Simulator app simulates data input into Eliona.

## Overview

This guide provides instructions on configuring, installing and using the Device Simulator app.

## Installation

Install the Device Simulator app via the Eliona App Store.

## Configuration

The Device Simulator app requires configuration through Eliona’s settings interface. Below are the general steps and details needed to configure the app.

### Configure the Device Simulator app 

Generators can be created in Eliona under `Apps > Device Simulator > Settings` which opens the app's [Generic Frontend](https://doc.eliona.io/collection/v/eliona-english/manuals/settings/apps). Here you can use the endpoint `/generators` with the POST method. Each generator requires the following data:

| Attribute         | Description                                                                     |
|-------------------|---------------------------------------------------------------------------------|
| `asset_id`        | ID of the asset being simulated.                                                |
| `attribute`       | Attribute name of the data being generated.                                     |
| `subtype`         | Subtype of the data being generated (e.g., "input").                             |
| `function_type`   | Type of function for data generation (e.g., "sin_wave", "random", "sawtooth_wave").|
| `min_value`       | Minimum value for the generated data.                                           |
| `max_value`       | Maximum value for the generated data.                                           |
| `integer`         | Specifies if the generated value should be an integer or float. (`true` for integer, `false` for float). |
| `interval_seconds`| Interval in seconds for data generation. (Sampling rate) (How often should the new data be generated?) |
| `frequency`       | Frequency in Hz for wave functions. (How often the function should reach the original value per second?) (typically less than 1) (Must not be higher than 1/(2*interval_seconds) to avoid aliasing / Nyquist frequency) |

Example generator JSON:

```json
{
  "asset_id": 17202,
  "attribute": "power",
  "subtype": "input",
  "function_type": "sin_wave",
  "min_value": 10.5,
  "max_value": 75.2,
  "integer": false,
  "interval_seconds": 1,
  "frequency": 0.05
}
```