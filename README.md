# Fioctl

A simple tool to interact with the Foundries.io REST API for managing a
Factory. Its based on the Foundries.io "ota-lite" API defined here:

[https://api.foundries.io/ota/](https://api.foundries.io/ota/)

## Using

You must first authenticate with the server before using this tool with:

~~~sh
fioctl login
~~~

Most commands require a "factory" argument. This can be defaulted inside
`$HOME/.config/fioctl.yaml`

~~~yaml
factory: <The name of your factory>
~~~

You can then view your fleet of devices with `fioctl device list`, or
start to see the Targets(ie "builds") applicable to your devices with the
`fioctl targets list`.

The rest of the commands can be discovered by running `fioctl device --help`
and `fioctl targets --help`.

## Building

### To build one binary for development

~~~sh
make build  # builds all targets

# or build for your specific target
make fioctl-linux-amd64
make fioctl-darwin-amd64
make fioctl-windows-amd64
~~~

### To build all binaries

**Prerequisite**
- install Gorelease: https://goreleaser.com/install/

Run:

```shell
goreleaser release --skip=publish --snapshot --clean
```

Check all binaries into the directory `dist/`

## Releasing

Just push a tag with the next version.

## Making Changes

After making changes be sure to run `make format` which will run the go-fmt
tool against the source code.
