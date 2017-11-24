# Log Generator

For log collector's performance test

# How to use


## Output fixed log lines

Write 100000000 logs at speed 100000 per seconds, this will ensure output 100000000 lines.

```
docker run -e MAX=100000000 -e SPEED=100000 \
  -e DEBUG=true -e OUTPUT=/tmp/abc/def.log \
  -v /host/path:/tmp liubin/loggenerator
```

## Run write test for fixed time duration

Or write logs at speed 100000 per seconds and continue 600 seconds, this can **NOT** ensure output 600000000 lines.

```
docker run -e DURATION=600 -e SPEED=100000 \
  -e DEBUG=true -e OUTPUT=/tmp/abc/def.log \
  -v /host/path:/tmp liubin/loggenerator
```
## Envs

* `MAX`: Total log size(linse) to write
* `SPEED`: Logs per second.
* `DEBUG`: If set, will print process info(speed, used time ...).
* `OUTPUT`: Log file full path(include log file name), if not provided, logs will be wrote to `stdout`
* `SIZE`: Size of one line
* `DURATION`: Write log continue time(seconds).
* `MODE`: `WEB` to use web api to controll, otherwise will be run as console cmd and use env as parameters.

`DURATION` and `MAX` must use one and can only use one.

## Web parameters

Only needed when `MODE=WEB`.



