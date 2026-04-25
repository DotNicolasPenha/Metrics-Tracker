# QUICK START

## 1 Download the Metracker

Select the your OS version in assets of this release: [v0.0.1-alpha.mt](https://github.com/DotNicolasPenha/metracker/releases/tag/v0.0.1-alpha.mt),
or select here:

### (darwin amd64 - mac): click [download](https://github.com/DotNicolasPenha/metracker/releases/download/v0.0.1-alpha.mt/mt.v0.0.1-alpha-darwin-amd64)
### (linux amd64 ): click [download](https://github.com/DotNicolasPenha/metracker/releases/download/v0.0.1-alpha.mt/mt.v0.0.1-alpha-linux-amd64)
### (windows exe): click [download](https://github.com/DotNicolasPenha/metracker/releases/download/v0.0.1-alpha.mt/mt.v0.0.1-alpha-windows-amd64.exe)

## 2 Create a interceptor with the save command:

Copy and run this command or learn about save cmd: [save.cmd.wiki.md](./cmds/save.cmd.wiki.md) 

```
mt save -n "ExampleInterceptor" --proxy-addr ":8080" --db-addr ":5432"
``` 

## 3 Run the interceptor:

Copy and run this command or learn about run cmd: [run.cmd.wiki.md](./cmds/run.cmd.wiki.md) 

```
mt run -n "ExampleInterceptor"
``` 

