# Repository structure:

- ./bin
Binare files.

- ./example/db  
Example database used for testing.

- ./build.sh  
Script to build the CLI.

- ./interceptor  
The core of Metracker. Responsible for creating and managing the database interceptor.

- ./user  
Handles persistence of user data in JSON format.

- ./cmd/root 
Configure root cobra cli.

- ./cmd/save 
Save cmd, to learn how to use in CLI, go to [save cmd wiki](./cmds/save.cmd.wiki.md).