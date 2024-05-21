### Salad Netdata Go plugin module

One of the Nedata-recommended ways to add custom charts to it is to extend its Go plugin my creating one or multiple modules. We will maintain the `salad` module in the Netdata repository fork.

#### Where the code is located in the repository?
The plugin source is located in `src/go/collectors/go.d.plugin`. Within that directory, the `modules/salad` subdirectory contains the Salad -specific code.

#### How to build it?
```
cd src/go/collectors/go.d.plugin
CGO_ENABLED=0 GOOS=linux go build -o go.d.plugin cmd/godplugin/main.go
```
This will produce the `go.d.plugin` binary to be eventually deployed to the server.

#### How to install it on the server?
- copy the binary to `/usr/libexec/netdata/plugins.d`
- Make sure the Go plugin is enabled: `/etc/netdata.conf` must contain lines
```
[plugins]
go.d = yes

```
The Go plugin configuration is located in `/etc/netdata/go.d.conf `. It describes which modules are enabled. the plugin ships with many, we don't need them so they can remain commented. Make sure that the `salad` module is enabled:
```
  salad: yes
```
This is a yaml file, so watch out the whitespace.

The module itself is configured in `etc/netdata/go.d/salad.conf`
For example, `update_every` indicates how often, in seconds, the module fetch the data from the running SGS server.

#### Smoke testing that everything is working?
Run
```
/usr/libexec/netdata/plugins.d/go.d.plugin -m salad
```
The output should not indicate any errors and look along these lines:
```

BEGIN 'salad.nodes'
SET 'active' = 5
SET 'quarantined' = 0
SET 'zombied' = 0
END

BEGIN 'netdata.execution_time_of_salad'
SET 'time' = 12
END

```
It will print more output, depending on the `update_every` value