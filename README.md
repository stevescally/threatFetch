# ThreatFetch

## Overview

Gather threat intel feeds from one or more sources.

## Functionality

- Flag for debugging
- Flag for logging
- Flag for disabling feed as source
- Create directory for feedname and sub-directories based on current date
- Include logic for specific feeds with config file override
  - I.e threatview.io only updates once a day at 11 UTC, defaults should support this
  - Configuration can allow over-ride

## Removed Functionality

- **Read configuration file for list of feed endpoints**
- **Present an index of available feeds read from configuration file and which are enabled.**

  The feeds are hard-coded and could require certain exceptions so a configuration at this
  point just seems like overhead. Since we are not reading a configuration file we don't 
  need the index idea.


## Layout

```
threatFetch
  |-- go.mod
  |-- threatFetch.go
  |-- threatviewio
      |-- threateviewio.go
  |-- <next feed> ..
```
