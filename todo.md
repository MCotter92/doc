# future features 

## Doc core functionality
    [] the user specifies a folder for notes (like vaults for obsidian)
    [] the user applies metadata to a file with frontmatter (tags have the option to be inline and notated with [[tag]])
    [] the metadata is collected by a database and stored
    [] the user can then search the database to find files
    [] things are automatically updated if the user makes a change to the metadata or moves a file

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
        [] delete 
        [] search 
        [] sync
    [] apply frontmatter parsing
    [] implement file syncing with database
    [] wrap in a rest api
        - python package fastapi
        - look into insomnia rest application version prior to 8
            - is there an equivelent for go? 
    [] look into test coverage
