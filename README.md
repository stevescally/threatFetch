# ThreatFetch

## Overview

Gather threat intel feeds from one or more sources.

## Functionality

- Read configuration file for list of feed endpoints
- Flag for debugging
- Flag for logging
- Create directory for feedname and sub-directories based on current date
- Include logic for specific feeds with config file override
  - I.e threatview.io only updates once a day at 11 UTC, defaults should support this
  - Configuration can allow over-ride
- Present an index of available feeds read from configuration file and which are enabled.
    - Might be a pterm option for this list
    - Update configuration with selections


## Layout

threatFetch
  |-- go.mod
  |-- threatFetch.go
  |-- threatviewio
      |-- threateviewio.go
  |-- <next feed> ..
