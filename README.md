# Is it alive?

A simple golang script made to verify if the websites listed in a file are up.

**Usage example:** 
`./isAlive urls.txt output.txt`

The script will only include URLs with a status code other than 404 (Not Found) in the output file.

*Note: An output file is optional; the script can run with just the URL list file.*


### Next update:
- add `goroutines` to check the status of multiple URLs simultaneously.
