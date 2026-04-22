# Save CMD Document

### The save command is the primary way to manage Metracker configurations. It follows an "Update or Create" logic: it attempts to load an existing configuration file, applies the changes provided via flags, and persists the updated state back to JSON.

Usage
Bash

./mt save [flags]

### Flags
--------------------------------------------------------
- -f, --file (Default: config.json): Path to the target configuration file.

- -m, --max-conn (Default: 100): Sets the maximum allowed active connections.

- -b, --block: Appends a new query string to the blacklist.

- -r, --retrys (Default: 3): Sets the retry limit for the query provided in the --block flag.

- --proxy-addr: Local address for the proxy to listen on (e.g., :5433).

- --db-addr: Remote database address to forward traffic to (e.g., localhost:5432).
---------------------------
### Examples
### 1. Initializing a Configuration

To create a new config.json with a specific connection limit:
```
 ./mt save --max-conn 150
``` 

### 2. Adding Security Rules

To block DROP TABLE queries and allow only 1 retry before action:
```
./mt save --block "DROP TABLE" --retrys 1
``` 
Note: This appends the rule to the existing list in the JSON file.
### 3. Setting Up an Interceptor

To map a local port to a destination database:
```
./mt save --proxy-addr ":5433" --db-addr "127.0.0.1:5432
```
### 4. Full Environment Setup

You can combine multiple flags to configure an entire environment in a single command:
``` 
./mt save -f prod.json -m 500 -b "DELETE FROM" -r 2 --proxy-addr ":6000" --db-addr "db.production.com:5432"
```
