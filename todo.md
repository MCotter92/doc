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
    [] apply frontmatter parsing
    [] implement core functionality 
        - create 
        - search 
        - delete 
        - update 
        - sync
    [] implement file syncing with database
    [] implement file watcher
        - thoughts on this file watcher. I think the only directory level watching would be if they moved thier notes folder. 
        - For file watching, I would want to watch for changes to the frontmatter on each file. 
        - If the user changes the keyword for a file, the file watcher triggers a sql query that updates the database accordingly.
        - I imagine the hard part will be updating the database, not having the file watcher shoot up a flair.
        - Maybe I can have a watcher spin up via a goroutine when a file in the notes directory is opened? And when the user saves, that goroutine also compares the current frontmatter to that file's entry in the database? If they are different, reconcile.  
    [] wrap in a rest api
        - python package fastapi
        - look into insomnia rest application version prior to 8
            - is there an equivelent for go? 
    [] look into test coverage
