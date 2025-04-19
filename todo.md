# future features 

## Doc core functionality
    [x] track a document 
    [x] create a document and track it
    [x] untrack and document
    [x] untrack and document and delete it
    [x] search for doucments based on metadata
    [] automatically update metadata if user updates the metadata outise of doc 

### overall ideas
    - have event tracker handle rm and whatnot and also have explicit commands   
        - if i delete a file with rm, but that file is still in doc's state, when i run doc it will recreate that file
            - if i hash the contents into the json it will return the contents too 

### todo list 
    - Struct population 
        [x] make sure last modified date is actually last modified date
        [x] make sure location has a control flow for if a document is not in current working directory. Need to parse string input from user for that if not "."
    [x] change input to this: doc search --title filename.md --keyword poop dumb --date mm/dd/yyyy
    [x] finish untrack document
    [x] refactor to implement factory design pattern. I want to make sure I'm using interfaces effectively.
    [] wrap in a rest api
        - python package fastapi
        - is there an equivelent for go? 
        - look into insomnia rest application version prior to 8
    [] apply mysql or sqlite
        - mysql workbench for my sql database 
            - start with tables for users and files
    [] implement a user profile 
        - should have the following
            [] editor of choice
            [] determin a notes folder to watch 
    [] implement tags by reading the file and looking for words surrounded by something (maybe __tag__)
    [] create a flag for cmd/track to track everything in a directory
    [] figure out how to add doc to my path and store the code somewhere other than dev like other programs.
        - give command to add to path, leave it up to them to run it. They can change that func as needed if they want. 
    [] look into test coverage
