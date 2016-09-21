# statsd-exec
`statsd-exec` is small shell command "wrapper" to run any arbitrary command and report execution time, success and failure information to [statsd](https://github.com/etsy/statsd).

> Note: Tested only on Linux and Mac. If you test on any other platform, please me know!

## Usage
```shell
statsd-exec <Add command or script here>
```

## Installation
- Download binary from [releases](https://github.com/ernoaapa/statsd-exec/releases)
- Give execution rights (`chmod +x statsd-exec`) and add it into your $PATH
- Add `STATSD_PREFIX="my.prefix" statsd-exec` in front of any command
- See below more examples!

### Configuration
Configuration happens through environment variables

| Environment variable | Required | Default   | Description                |
|----------------------|----------|-----------|----------------------------|
| STATSD_METRIC_NAME   | true     | -         | Name of the metric         |
| STATSD_PREFIX        | false    | stats     | Prefix for all stats       |
| STATSD_HOST          | false    | localhost | Hostname for Statsd client |
| STATSD_PORT          | false    | 8125      | Port for Statsd client     |

### Examples
#### Successful execution
If command returns successfully (exit code 0)
```
STATSD_METRIC_NAME="process.helloworld" statsd-exec echo "Hello World"
```

This would send following metrics:
- `stats.process.helloworld.executed:1|c`
- `stats.process.helloworld.duration:8.167651|ms`
- `stats.process.helloworld.success:1|c`
- `stats.process.helloworld.failed:0|c`

#### Failure execution
If command returns failure
```
STATSD_METRIC_NAME="process.helloworld" statsd-exec test "foo" == "bar"
```

This would send following metrics:
- `stats.process.helloworld.executed:1|c`
- `stats.process.helloworld.duration:6.332267|ms`
- `stats.process.helloworld.success:0|c`
- `stats.process.helloworld.failed:1|c`

#### Custom prefix
You can for instance set global prefix for all metrics and define only metric name for each execution
```
export STATSD_PREFIX="stats.testing"
STATSD_METRIC_NAME="foo" statsd-exec echo "Hello Foo"
STATSD_METRIC_NAME="bar" statsd-exec echo "Hello Bar"
```

This would send following metrics:
- `stats.testing.foo.*`
- `stats.testing.bar.*`

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
