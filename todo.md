# future features 

## mk1 
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
[] implement tags by reading the file and looking for words surrounded by something (maybe __tag__)
[] implement event listener or an audit function that the user manually calls. 
    - autoupdate inline tags.
    - either way it should do the following: 
    - update file structs dynamically 
    - delete files removed by user with rm 
[] create a flag for cmd/track to track everything in a directory
[] figure out how to add doc to my path and store the code somewhere other than dev like other programs.
    - give command to add to path, leave it up to them to run it. They can change that func as needed if they want. 
[] look into test coverage
## mk2 
[] look into how to hash file contents into json
    - have event tracker handle rm and whatnot and also have explicit commands   
        - if i delete a file with rm, but that file is still in doc's state, when i run doc it will recreate that file
            - if i hash the contents into the json it will return the contents too 
[] explore frontmatter
[] read about terriform, terriform state, terriform state surgery  
