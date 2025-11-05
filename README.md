# Doc
A CLI for organizing your documents. 


## Doc core functionality for 1.0

- the user specifies a folder for notes (like vaults for obsidian)
- doc applies metadata to a file with frontmatter that is that file's sql entry data
- the metadata is collected by a database and stored
- the user can then search the database to find files
- things are automatically updated in the database if the user makes a change to the metadata or moves a file

## features for after 1.0

- maybe a cool markdown realtime rendering thing like in obsidian? so the user is looking at nice notes and not hashtags everywhere? 
    - this seems hard.
- support for cloud storage and team colaboration essentiall making this a SWE team documentation platform. I don't need this but it would be a good vehicle for learning networking and devops stuff (aws, terriform, etc). Would look good on a resume.


### **TODO LIST**

## Done

- [x] implement a doc init command that applies a .doc folder with a yaml file inside for user config
- [x] implement yaml config file
    - should have the following
    - [x] editor of choice
    - [x] determin a notes folder to watch 
    - [x] implement a doc config command that cats out the user's config file
- [x] create 
- [x] open - search then open
- [x] delete - search then update 

## Now
**Make it work, then make it good.**

- [] make sure that UpdateCriteria and SearchCriteria structs match up with Doc struct in naming. 
- [] make sure cli flags match up with Doc struct.
    - do I need directory in doc struct? 
- [] update - maybe if there are no flags, then show a table output where users can find what they want. or if the user provies
the full path, then they can say -k=newKeyword to update the keyword.
    - update db
        - [] make db.Update() func. 
        - [] make docCore.Update() 
            - [] calls db.Search then db.Update 

    - update frontmatter
        - [] make updateFrontmatter() 
        - [] make unmarshalFrontmatter(pathToFile) *someStruct {}
            - needs to 
                - 1. read file and marshal yaml into a struct
                - 2. use updateCriteria fields to edit the marshaled yaml
                - 3. overwrites the yaml in the file and only the yaml
        - [] config 
            - this is just waiting on updateUserConfig for now.

## Next 
- [] doc sync
    - probably search then update en masse? 
- [] implement REST API 
- [] go back and refactor cmd/ to utilize integration points effectively. 
    - https://cobra.dev/docs/explanations/philosophy
- [] apply file watcher for frontmatter parsing
    - if the user changes the frontmatter in a file, the file watcher should trigger the update command on that file
- [] apply inline backlinks 
    - should these be part of the db? 
    - user could say give me all my notes tagged to this note. Just a thought. Idk how this would change the db schema.

## Later 
- [] doc manage 
    - maybe allow user to do whatever they need from here. select the delete or open etc. maybe a tui that ties all my CRUD together?
    - eventually allows for management of AWS stuff too 
- [] implement automatic file syncing with database? 
- [] implement some AWS functionality like Jake suggested a LONG time ago. Dynamo DB backupd w/ IaC? Something fun to think about
- [] implement unit testing for easier refactoring 
- [] can i trim down my imports? aquasecurity/table might not be really needed tbh. I bet I can figure that out.

