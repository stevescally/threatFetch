# ThreatFetch

## Overview

Gather "threat"" intel feeds from multiple publicly available sources.

## Functionality

### Implemented

###### Download Directory Path

Currently the (-p) flag is available to specify the root directory where the feeds should
download to. This can be useful for doing a one-time run test or downloading the latest 
feed data for analysis.

The *Default* option is to download to the current working directory.

###### Feed Sources
- [threatview.io](https://threatview.io/) (8) feeds
  - osintThreatFeed
  - c2HuntFeed
  - IPBlocklist
  - domainBlocklist
  - MD5HashBlocklist
  - URLBlocklist
  - bitcoinAddressIntel
  - SHAFilehashBlocklist
- [greensnow.co](https://www.greensnow.co/) (1) feed
  - IPBlocklist

### Planned

###### Flag for debugging

Implement debug flag (-d) which will use the pterm logger to print existing pterm.debug output.

###### Flag for logging

Implement logging flag (-l) that will write output to a log file.

###### Flag for disabling feed as source

Implement disabling a feed source flag (-dfs). I have seen cli tools that can do this by name or
by number. User can run status flag to get the feed ordering. 

  examples:
  ```
  -dfs 1,3,5
  -dfs 1-5
  -dfs threatview, greensnow
 ````

###### Flag for feed status table

Implement a status flag (-s) providing the following information for a feed source:

|Number|Feed Name|Directory Path|Last Update|Operational Status|Feed Site|
|------|---------|--------------|-----------|------------------|---------|
|1|threatview|threatview_data|20240213|Online|https://threatview.io|

### Ideating 

- Create directory for feedname and sub-directories based on current date
  - Would this be useful if the file is already dated?
- Read configuration file for list of feed endpoints
  - This might work if the structure was all the same.
  - Feeds are already added after reviewing the list so value is minimal.
- Present an index of available feeds read from configuration file and which are enabled.
  - Similar to the reading of a configuration file this doesn't seem useful.
  - Maybe status of feed sources based on currently downloads would be useful.

  The feeds are hard-coded and could require certain exceptions so a configuration at this
  point just seems like overhead. Since we are not reading a configuration file we don't 
  need the index idea.

- Include logic for specific feeds with config file override
  - I.e threatview.io only updates once a day at 11 UTC, defaults should support this
  - Configuration can allow over-ride

## Layout

Currently some of the functions are repeated per source with only the structs being changed.
I will re-evaluate this as feeds are added as I'm not concerned with code being repeated as
much as I am trying to figure out complex logic to account for how feeds present their data.
```
threatFetch
  |-- go.mod
  |-- threatFetch.go
  |-- threatviewio
      |-- threateviewio.go
  |-- <next feed> ..
```

## License

###### Code (CC BY-NC-SA)

Reference the [Creative Commons](https://creativecommons.org/) site for the current [BY-NC-SA-4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode.txt) text. 

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.

###### Feed Data

The feeds currently implemented use only publicly available sources and where indicated,
to the best of any ability, the proper abuse limitations have been added, such as checking 
if the feed has previously been downloaded. 

Assume the feeds are for non-commercial use, unless explicitly indicated
otherwise by referencing the individual feed site for their terms and
conditions.  
