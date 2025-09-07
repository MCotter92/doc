# Doc
A CLI for organizing your documents inspired by Obsidian.

**Make it work, then make it good.**

## Doc core functionality for 1.0
    - the user specifies a folder for notes (like vaults for obsidian)
    - doc applies metadata to a file with frontmatter that is that file's sql entry data
    - the metadata is collected by a database and stored
    - the user can then search the database to find files
    - things are automatically updated in the database if the user makes a change to the metadata or moves a file

## features for after 1.0
    - maybe a cool markdown realtime rendering thing like in obsidian? so the user is looking at nice notes and not hashtags everywhere? 
    - support for cloud storage and team colaboration essentiall making this a SWE team documentation platform. I don't need this but it would be a good vehicle for learning networking and devops stuff (aws, terriform, etc). Would look good on a resume.


### todo list 
    [x] implement a doc init command that applies a .doc folder with a yaml file inside for user config
    [x] implement yaml config file
        - should have the following
            [x] editor of choice
            [x] determin a notes folder to watch 
    [x] implement a doc config command that cats out the user's config file
    [x] reorganize utils 
    [] implement core functionality 
        [x] create 
        [x] open - search then open
        [] delete - search then update 
        [] sync - probably search then update en masse? 
        later - [] explore - not just internal search used, might want a search command available to the user so they can explore.
            - i might want to do this later. I think a cool TUI interface would be good. Something like telescope? Sounds like a big project
    [] apply inline backlinks 
        - should these be part of the db? 
        - user could say give me all my notes tagged to this note. Just a thought. Idk how this would change the db schema.
    [] apply file watcher for frontmatter parsing
        - if the user changes the frontmatter in a file, the file watcher should trigger the update command on that file
    [] implement automatic file syncing with database? 
    [] wrap in a rest api
    [] implement unit testing for easier refactoring 

