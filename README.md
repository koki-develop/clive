<p align='center'>
<img src="./assets/logo_light.svg#gh-light-mode-only" />
<img src="./assets/logo_dark.svg#gh-dark-mode-only" />
</p>

<p align="center">
Automates terminal operations and lets you view them live via a browser.
</p>

<p align='center'>
<a href="https://github.com/koki-develop/clive/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/clive?style=flat" /></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/clive?style=flat" /></a>
<a href="https://github.com/koki-develop/clive/actions/workflows/ci.yml"><img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/koki-develop/clive/ci?logo=github&style=flat" /></a>
<a href="https://codeclimate.com/github/koki-develop/clive/maintainability"><img alt="Code Climate maintainability" src="https://img.shields.io/codeclimate/maintainability/koki-develop/clive?logo=codeclimate&style=flat" /></a>
</p>

![](./examples/demo/demo.gif)

# cLive

- [Prerequisite](#prerequisite)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Commands](#commands)
- [Configuration](#configuration)
  - [Actions](#actions)
  - [Settings](#settings)
- [Examples](#examples)
- [License](#license)

## Prerequisite

cLive requires [ttyd](https://tsl0922.github.io/ttyd/) to be installed.  
For example, if you use homebrew, you can install it with `brew install ttyd`.

```sh
$ brew install ttyd
```

See [ttyd documentation](https://github.com/tsl0922/ttyd#installation) for more information.

## Installation

> **Note**
> There are prerequisites for using cLive. See [`Prerequisite`](#prerequisite) for details.

cLive can be installed with `go install`.

```sh
$ go install github.com/koki-develop/clive@latest
```

Or download the binary from the [releases page](https://github.com/koki-develop/clive/releases/latest).

## Getting Started

First, run `clive init`.

```sh
$ clive init
Created ./clive.yml
```

A file named `clive.yml` will then be created with the following contents:

```yaml
# documentation: https://github.com/koki-develop/clive#settings
settings:
  loginCommand: ["bash", "--login"]
  fontSize: 22
  defaultSpeed: 10

# documentation: https://github.com/koki-develop/clive#actions
actions:
  - pause
  - type: echo 'Welcome to cLive!'
  - key: enter
```

Finally, run `clive start` to launch the browser and start cLive.

```sh
$ clive start
```

## Commands

Available commands:

- [`init`](#clive-init) - Create config file.
- [`start`](#clive-start) - Load config file and start cLive.
- [`completion`](#clive-completion) - Generate the autocompletion script for the specified shell.

### `clive init`

Create a config file.

```sh
$ clive init
```

| Flag | Default | Description |
| --- | --- | --- |
| `-c`, `--config` | `./clive.yml` | Config file name. |

### `clive start`

Load config file and start cLive.  
See [`Configuration`](#configuration) for config file.

```sh
$ clive start
```

| Flag | Default | Description |
| --- | --- | --- |
| `-c`, `--config` | `./clive.yml` | Config file name. |

### `clive completion`

Generate the autocompletion script for clive for the specified shell.  
See each sub-command's help for details on how to use the generated script.

```sh
$ clive completion <shell>

# e.g.
$ clive completion bash
$ clive completion bash --help
```

Available shells:

- bash
- fish
- powershell
- zsh

## Configuration

Config file consists of `actions` and `settings`.

- [`actions`](#actions) - Actions to run.
- [`settings`](#settings) - Basic settings (font size, default speed, etc.).

### Actions

Actions to run.  
Available actions:

- [`type`](#type) - Type characters.
- [`key`](#key) - Enter special keys.
- [`ctrl`](#ctrl) - Enter the ctrl key with other keys.
- [`sleep`](#sleep) - Sleep for a specific number of milliseconds.
- [`pause`](#pause) - Pause actions.

#### `type`

Type characters.

| Field | Required | Default | Description |
| --- | --- | --- | --- |
| `type` | **Yes** | N/A | Characters to type. |
| `count` | No | `1` | Number of times the action is repeated. |
| `speed` | No | `10` | Interval between key typing (milliseconds). |

```yaml
# e.g.
actions:
  - type: echo 'Hello World'
    count: 10
    speed: 100
```

#### `key`

Enter special keys.  
Available keys:

- `esc`
- `backspace`
- `tab`
- `enter`
- `left`
- `up`
- `right`
- `down`

| Field | Required | Default | Description |
| --- | --- | --- | --- |
| `key` | **Yes** | N/A | Special key to type. |
| `count` | No | `1` | Number of times the action is repeated. |
| `speed` | No | `10` | Interval between key typing (milliseconds). |

```yaml
# e.g.
actions:
  - key: enter
    count: 10
    speed: 100
```

#### `ctrl`

Enter the ctrl key with other characters.

| Field | Required | Default | Description |
| --- | --- | --- | --- |
| `ctrl` | **Yes** | N/A | Characters to enter with the ctrl key. |
| `count` | No | `1` | Number of times the action is repeated. |
| `speed` | No | `10` | Interval between key typing (milliseconds). |

```yaml
# e.g.
actions:
  - ctrl: c
    count: 10
    speed: 100
```

#### `sleep`

Sleep for a specific number of milliseconds.

| Field | Required | Default | Description |
| --- | --- | --- | --- |
| `sleep` | **Yes** | N/A | Time to sleep (milliseconds). |

```yaml
# e.g.
actions:
  - sleep: 3000 # Sleep for 3 seconds.
```

#### `pause`

Pause actions.  
Press enter to continue.

```yaml
# e.g.
actions:
  - pause
```

### Settings

Basic settings.  
Available settings:

- [`loginCommand`](#logincommand) - Login command and args.
- [`fontSize`](#fontsize) - Font size.
- [`fontFamily`](#fontfamily) - Font family
- [`defaultSpeed`](#defaultspeed) - Default interval between key typing.
- [`browserBin`](#browserbin) - Path to executable browser binary.

#### `loginCommand`

Set command and args for logging into the shell.  
Default: `["bash", "--login"]`.

```yaml
# e.g.
settings:
  loginCommand: ["zsh", "--login"]
```

#### `fontSize`

Set font size.  
Default: `22`

```yaml
# e.g.
settings:
  fontSize: 36
```

#### `fontFamily`

Set font family.  

```yaml
# e.g.
settings:
  fontFamily: monospace
```

#### `defaultSpeed`

Set default interval between key typing (milliseconds).  
Default: `10`

```yaml
# e.g.
settings:
  defaultSpeed: 100
```

#### `browserBin`

Set path to executable binary for the browser used.  
See [go-rod documentation](https://github.com/go-rod/go-rod.github.io/blob/master/compatibility.md#supported-browsers) for supported browsers.

```yaml
# e.g.
settings:
  browserBin: /Applications/Sidekick.app/Contents/MacOS/Sidekick # use Sidekick
```

## Examples

For more examples see the [`examples/`](./examples/) directory.

## License

[MIT License](./LICENSE)
