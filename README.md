A CLI for organizing your documents inspired by Obsidian.

# Doc core functionality
    - the user specifies a folder for notes (like vaults for obsidian)
    - the user applies metadata to a file with frontmatter (tags (or backlinks) should be inline and notated with [[tag]])
    - the metadata is collected by a database and stored
    - the user can then search the database to find files
    - things are automatically updated if the user makes a change to the metadata or moves a file

## todo list 
    [x] implement a doc init command that applies a .doc folder with a yaml file inside for user config
    [x] implement yaml config file
        - should have the following
            [x] editor of choice
            [x] determin a notes folder to watch 
    [x] implement a doc config command that cats out the user's config file
    [x] reorganize utils 
    [] implement core functionality 
        [x] create 
        [x] search 
        [] delete - search then update 
        [] update - search then update
        [] sync - probably search then update en masse? 
    [] apply inline backlinks 
        - should these be part of the db? 
        - user could say give me all my notes tagged to this note. Just a thought. Idk how this would change the db schema.
    [] apply file watcher
    [] apply frontmatter parsing
        - if the user changes the frontmatter in a file, the file watcher should trigger the update command on that file
    [] implement automatic file syncing with database? - is this necessary given the sync command? 
    [] wrap in a rest api
    [] look into test coverage

