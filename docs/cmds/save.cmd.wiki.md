# Save Command Documentation

The `save` command is the primary tool for managing Metracker's interceptors. Following an **"Update or Create"** logic, it allows you to target specific interceptors by name to modify their security rules and connection limits.

## Usage

```bash
./mt save [flags]
```

## Flags

| Short | Long          | Default       | Description                                           |
| :---- | :------------ | :------------ | :---------------------------------------------------- |
| `-f`  | `--file`          | `config.json` | Path to the target configuration file.               |
| `-n`  | `--name`          | `default`     | **Crucial:** The unique name of the interceptor to manage. |
| `-m`  | `--max-conn`      | `100`         | Maximum active connections for this specific interceptor. |
| `-b`  | `--block`         | `""`          | Appends a query string to this interceptor's blacklist. |
| `-r`  | `--retrys`        | `3`           | Retry limit for the query provided in the `--block` flag. |
|       | `--proxy-addr`    | `""`          | Local address for the proxy to listen on.             |
|       | `--db-addr`       | `""`          | Remote database address to forward traffic to.        |
| `-i`  | `--authorized-ips`| `""`      | Appends a ip to ips authorized list of interceptor        |

---

## Examples

### 1. Creating a New Named Interceptor
To set up a specific mapping for a Postgres instance:
```bash
./mt save --name "postgres-prod" --proxy-addr ":5433" --db-addr "localhost:5432"
```

### 2. Updating Limits for a Specific Interceptor
If you have multiple interceptors, use the `--name` flag to target the correct one:
```bash
./mt save --name "postgres-prod" --max-conn 250
```

### 3. Adding Security Rules Per Interceptor
Add a blacklist rule only to the "billing-db" interceptor:
```bash
./mt save -n "billing-db" -b "DROP TABLE" -r 1
```

### 4. Full Environment Setup
Configure a complete named interceptor in a single line:
```bash
./mt save -f services.json -n "auth-db" -m 500 -b "DELETE FROM" -r 2 --proxy-addr ":6000" --db-addr "auth.internal:5432"
```

---