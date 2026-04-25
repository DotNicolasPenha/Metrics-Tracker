# QUICK START

ALERT: THIS IS AN EXPERIMENTAL PROJECT.
## 1. Download Metracker

Select the version for your OS from the assets of this release: v0.0.1-alpha.mt, or choose below:
- (darwin amd64 - mac): click [download](https://github.com/DotNicolasPenha/metracker/releases/download/v0.0.1-alpha.mt/mt.v0.0.1-alpha-darwin-amd64)
- (linux amd64 ): click [download](https://github.com/DotNicolasPenha/metracker/releases/download/v0.0.1-alpha.mt/mt.v0.0.1-alpha-linux-amd64)
- (windows exe): click [download](https://github.com/DotNicolasPenha/metracker/releases/download/v0.0.1-alpha.mt/mt.v0.0.1-alpha-windows-amd64.exe)

(Rename the binary to 'mt')
## 2. Create an Interceptor

Run the following command to create an interceptor, or learn more about the save command here: [save.cmd.wiki.md](./cmds/save.cmd.wiki.md) 

``` 
./mt save -n "ExampleInterceptor" --proxy-addr ":8080" --db-addr "localhost:5432"
```
(Note: --db-addr should be your database address).
Update your application's connection:

Point your application to the --proxy-addr instead of the direct database address:
Go
```
// env:
// :8080 is your proxy address
// example:
connstring = "postgresql://admin:1234@localhost:8080/databaseOfMyApplication"
```
## 3. Run the Interceptor

Run the command below, or learn more about the run command here: [run.cmd.wiki.md](./cmds/run.cmd.wiki.md) 

```
./mt run -n "ExampleInterceptor"
```
You can check the examples directory if you prefer to see it in action without a manual setup.
