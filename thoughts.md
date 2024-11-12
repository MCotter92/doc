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
## Doc make (In prog) 
- Makes a new document. 
  - Add the option to implement tracking now. Optional params? 
## Doc open 
- a way to open one or a set of docs 
- For all, for any open, or just list them 
## Metadata stored in JSON 
- Individual files - slightly more space 
- State files - slower to query 
Will probably end up with both
Storing these in a document database that ships with doc 
This becomes the parent json 
Can query entire set of metadata files to get what i want 
By tag, date, keyword etc 
One for each child namespace 
Global namespace that manages everything 
This namespace knows where docs are to limit search 
Can add, delete, or update a namespace for each child location 
Knows what tags, subtags exist 

Go 
Go interfaces are weird
Relationships aren’t like “this has that” more like “this does that” 

Need to read about interfaces. 
