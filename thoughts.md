# Doc - a document org cli 

- Compatible with txt and md files for starters (seems like its all files right now?)

## Doc search 
- date or date range 
- title
- keywords
- Is Json the best way to do this? 
## Doc track 
- Add a document to doc’s tracking
- When a document is deleted, should doc automatically delete that metadata out of the json? 
- How do I tackle this? 
  - Well so it would look like "doc track path/to/doc.md", doc would parse out the necessary info for title, location.
  Need to make sure that today's day isn't being automatically populated, and the user can add keywords if they want. 
  - Lots of overlap with "doc make"" here... 
  - Interfaces??
## Doc make (In prog) 
- Makes a new document. 
  - Add the option to implement tracking now. Optional params? 
## Doc open 
- a way to open one or a set of docs 
- For all, for any open, or just list them 
## Metadata stored in JSON 
- Individual files - slightly more space 
- State files - slower to query 

Go interfaces are weird
Relationships aren’t like “this has that” more like “this does that” 

Need to read about interfaces. 
