# example section

## example.go 
A simple example for using the csv2json package
## addfields.go 
This example adds a bit more complexity by adding additional fields to the output json, which are not in the original csv file.
In this example the file contains a name and a date, these are parsed and used as additional fields.
## options.go
Shows the usage of options
- LineWiseJson: 
 - Since we needed this feature for our ELK stack this implements that every csv line equals a valid json line. The resulting file is not valid json but contains valid json for every line !
- QuoteEverything:
 - ignores datatypes and quotes everything
- PrettyPrint:
 - generates a more readable output (multilines)
