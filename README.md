# statsd-exec
`statsd-exec` is small shell command "wrapper" to run any arbitrary command and report execution time, success and failure information to [statsd](https://github.com/etsy/statsd).

## Usage
```shell
statsd-exec <Add command or script here>
```

### Configuration
Configuration happens through environment variables

| Environment variable | Required | Default   | Description                |
|----------------------|----------|-----------|----------------------------|
| STATSD_PREFIX        | true     | -         | Prefix for all stats       |
| STATSD_HOST          | false    | localhost | Hostname for Statsd client |
| STATSD_PORT          | false    | 8125      | Port for Statsd client     |

### Examples
#### Successful execution
If command returns successfully (exit code 0)
```
STATSD_PREFIX="process.helloworld" statsd-exec echo "Hello World"
```

This would send following metrics:
- `process.helloworld.executed:1|c`
- `process.helloworld.duration:8.167651|ms`
- `process.helloworld.success:1|c`
- `process.helloworld.failed:0|c`

#### Failure execution
If command returns failure
```
STATSD_PREFIX="process.helloworld" statsd-exec test "foo" == "bar"
```

This would send following metrics:
- `process.helloworld.executed:1|c`
- `process.helloworld.duration:6.332267|ms`
- `process.helloworld.success:0|c`
- `process.helloworld.failed:1|c`


## Development
### Get dependencies
```shell
go get ./...
```

### Run Statsd locally
See [statsd](https://github.com/etsy/statsd) documentation how to run it locally.
Highly suggest to set `debug:true` and `dumpMessages:true` in the configuration to see the messages in terminal output.

### Run stastd-exec
```shell
STATSD_PREFIX="testing" go run main.go echo "Hello World!"
```
