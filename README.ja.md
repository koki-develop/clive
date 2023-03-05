<p align='center'>
<img src="./assets/logo_light.svg#gh-light-mode-only" />
<img src="./assets/logo_dark.svg#gh-dark-mode-only" />
</p>

<p align="center">
Automates terminal operations and lets you view them live via a browser.
</p>

<p align='center'>
<a href="https://github.com/koki-develop/clive/releases/latest"><img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/koki-develop/clive?style=flat"></a>
<a href="https://github.com/koki-develop/clive/releases/latest"><img alt="GitHub all releases" src="https://img.shields.io/github/downloads/koki-develop/clive/total?style=flat"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/clive?style=flat" /></a>
<a href="https://github.com/koki-develop/clive/actions/workflows/ci.yml"><img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/koki-develop/clive/ci.yml?branch=main&logo=github&style=flat" /></a>
<a href="https://codeclimate.com/github/koki-develop/clive/maintainability"><img alt="Code Climate maintainability" src="https://img.shields.io/codeclimate/maintainability/koki-develop/clive?logo=codeclimate&style=flat" /></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/clive"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/koki-develop/clive" /></a>
<a href="https://codecov.io/github/koki-develop/clive"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/koki-develop/clive?logo=codecov" /></a>
</p>

![](./examples/demo/demo.gif)

<p align="center">
<a href="./README.md">English</a> | 日本語
</p>

# cLive

- [前提条件](#information_source-前提条件)
- [インストール](#zap-インストール)
- [基本的な使い方](#beginner-基本的な使い方)
- [コマンド](#rocket-コマンド)
  - [`init`](#clive-init)
  - [`start`](#clive-start)
  - [`validate`](#clive-validate)
  - [`completion`](#clive-completion)
- [設定](#gear-設定)
  - [`actions`](#actions)
  - [`settings`](#settings)
- [サンプル](#book-サンプル)
- [ライセンス](#memo-ライセンス)

## :information_source: 前提条件

cLive は事前に [ttyd](https://tsl0922.github.io/ttyd/) がインストールされている必要があります。  
例えば homebrew を使用している場合、 `brew install` でインストールすることができます。

```sh
$ brew install ttyd
```

それ以外のインストール方法については [ttyd のドキュメント](https://github.com/tsl0922/ttyd#installation)を参照してください。

## :zap: インストール

> **Note**
> cLive を使用するには前提条件があります。詳しくは [`前提条件`](#information_source-前提条件) を参照してください。

もし homebrew を使用している場合、 `brew install` でインストールすることができます。

```sh
$ brew install koki-develop/tap/clive
```

また、 `go install` でもインストールすることができます。

```sh
$ go install github.com/koki-develop/clive@latest
```

もしくは [Releases ページ](https://github.com/koki-develop/clive/releases/latest)からバイナリをダウンロードしてください。

## :beginner: 基本的な使い方

まず `clive init` を実行します。

```sh
$ clive init
Created ./clive.yml
```

すると、 `clive.yml` という名前で次のような内容のファイルが作成されます。

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

最後に `clive start` を実行するとブラウザが立ち上がり、 cLive が開始されます。

```sh
$ clive start
```

## :rocket: コマンド

- [`init`](#clive-init) - 設定ファイルを作成します。
- [`start`](#clive-start) - 設定ファイルを読み込み、 cLive を開始します。
- [`completion`](#clive-completion) - 指定されたシェルの自動補完スクリプトを生成します。

### `clive init`

設定ファイルを作成します。

```sh
$ clive init
```

| フラグ           | デフォルト    | 説明             |
| ---------------- | ------------- | ---------------- |
| `-c`, `--config` | `./clive.yml` | 設定ファイル名。 |

### `clive start`

設定ファイルを読み込み、 cLive を開始します。
設定ファイルについては[`設定`](#gear-設定)を参照してください。

```sh
$ clive start
```

| フラグ           | デフォルト    | 説明             |
| ---------------- | ------------- | ---------------- |
| `-c`, `--config` | `./clive.yml` | 設定ファイル名。 |

### `clive validate`

設定ファイルを検証します。

```sh
$ clive validate
```

| Flag             | Default       | Description      |
| ---------------- | ------------- | ---------------- |
| `-c`, `--config` | `./clive.yml` | 設定ファイル名。 |

### `clive completion`

指定されたシェルの自動補完スクリプトを生成します。  
生成されたスクリプトの使い方についてはヘルプを参照してください。

```sh
$ clive completion <shell>

# e.g.
$ clive completion bash
$ clive completion bash --help
```

サポートしているシェル:

- bash
- fish
- powershell
- zsh

## :gear: 設定

設定ファイルは `actions` と `settings` で構成されます。

- [`actions`](#actions) - 実行するアクションのリスト。
- [`settings`](#settings) - 基本的な設定 (フォントサイズ、デフォルトの速度など) 。

### `actions`

実行するアクションです。  
有効なアクション:

- [`type`](#type) - 文字を入力します。
- [`key`](#key) - 特殊キーを入力します。
- [`ctrl`](#ctrl) - Ctrl キーを他のキーと一緒に入力します。
- [`sleep`](#sleep) - 指定した時間スリープします。
- [`pause`](#pause) - アクションを一時停止します。
- [`screenshot`](#screenshot) - スクリーンショットを撮影します。

#### `type`

文字を入力します。

| フィールド | 必須    | デフォルト | 説明                           |
| ---------- | ------- | ---------- | ------------------------------ |
| `type`     | **Yes** | N/A        | 入力する文字。                 |
| `count`    | No      | `1`        | アクションを繰り返す回数。     |
| `speed`    | No      | `10`       | キーを入力する間隔 (ミリ秒) 。 |

```yaml
# e.g.
actions:
  - type: echo 'Hello World'
    count: 10 # Optional
    speed: 100 # Optional
```

#### `key`

特殊キーを入力します。  
使用できるキー:

- `esc`
- `backspace`
- `tab`
- `enter`
- `left`
- `up`
- `right`
- `down`
- `space`

| フィールド | 必須    | デフォルト | 説明                           |
| ---------- | ------- | ---------- | ------------------------------ |
| `key`      | **Yes** | N/A        | 入力するキー。                 |
| `count`    | No      | `1`        | アクションを繰り返す回数。     |
| `speed`    | No      | `10`       | キーを入力する間隔 (ミリ秒) 。 |

```yaml
# e.g.
actions:
  - key: enter
    count: 10 # Optional
    speed: 100 # Optional
```

#### `ctrl`

Ctrl キーを他のキーと一緒に入力する。

| フィールド | 必須    | デフォルト | 説明                            |
| ---------- | ------- | ---------- | ------------------------------- |
| `ctrl`     | **Yes** | N/A        | Ctrl キーと一緒に入力するキー。 |
| `count`    | No      | `1`        | アクションを繰り返す回数。      |
| `speed`    | No      | `10`       | キーを入力する間隔 (ミリ秒) 。  |

```yaml
# e.g.
actions:
  - ctrl: c # Ctrl+c
    count: 10 # Optional
    speed: 100 # Optional
```

#### `sleep`

指定した時間スリープします。

| フィールド | 必須    | デフォルト | 説明                         |
| ---------- | ------- | ---------- | ---------------------------- |
| `sleep`    | **Yes** | N/A        | スリープする時間 (ミリ秒) 。 |

```yaml
# e.g.
actions:
  - sleep: 3000 # 3 秒間スリープする
```

#### `pause`

アクションを一時停止します。  
エンターキーを入力して再開します。

```yaml
# e.g.
actions:
  - pause
```

#### `screenshot`

スクリーンショットを撮影します。  
撮影したスクリーンショットは `screenshots/` に保存されます.  
スクリーンショットを保存するディレクトリは [`settings.screenshotsDir`](#screenshotsdir) で変更することができます。

```yaml
# e.g.
actions:
  - screenshot
```

### `settings`

基本的な設定です。  
設定できる項目:

- [`loginCommand`](#logincommand) - ログインコマンドと引数。
- [`fontSize`](#fontsize) - フォントサイズ。
- [`fontFamily`](#fontfamily) - フォントファミリー。
- [`defaultSpeed`](#defaultspeed) - デフォルトの入力速度。
- [`skipPauseBeforeQuit`](#skippausebeforequit) - 終了前の一時停止をスキップするかどうか。
- [`screenshotsDir`](#screenshotsdir) - スクリーンショットを保存するディレクトリ。
- [`browserBin`](#browserbin) - ブラウザの実行可能なバイナリへのパス。
- [`headless`](#headless) - ブラウザをヘッドレスモードで動作させるかどうか。
- [`width`](#width) - ウィンドウの幅。
- [`height`](#height) - ウィンドウの高さ。

#### `loginCommand`

シェルへのログインコマンドと引数を設定します。  
デフォルト: `["bash", "--login"]`.

```yaml
# e.g.
settings:
  loginCommand: ["zsh", "--login"]
```

#### `fontSize`

フォントサイズを設定します。  
デフォルト: `22`

```yaml
# e.g.
settings:
  fontSize: 36
```

#### `fontFamily`

フォントファミリーを設定します。

```yaml
# e.g.
settings:
  fontFamily: monospace
```

#### `defaultSpeed`

デフォルトの入力速度を設定します (ミリ秒) 。  
デフォルト: `10`

```yaml
# e.g.
settings:
  defaultSpeed: 100
```

#### `skipPauseBeforeQuit`

終了前の一時停止をスキップするかどうか。  
デフォルト: `false`

```yaml
# e.g.
settings:
  skipPauseBeforeQuit: true
```

#### `screenshotsDir`

スクリーンショットを保存するディレクトリ。  
デフォルト: `screenshots`

```yaml
# e.g.
settings:
  screenshotsDir: path/to/pictures
```

#### `browserBin`

ブラウザの実行可能なバイナリへのパスを設定します。  
サポートしているブラウザについては [go-rod のドキュメント](https://github.com/go-rod/go-rod.github.io/blob/master/compatibility.md#supported-browsers) を参照してください。

```yaml
# e.g.
settings:
  browserBin: /Applications/Sidekick.app/Contents/MacOS/Sidekick # Sidekick を使う
```

#### `headless`

ブラウザをヘッドレスモードで動作させるかどうか。  
デフォルト: `false`

```yaml
# e.g.
settings:
  headless: true
```

#### `width`

ウィンドウの幅。

```yaml
# e.g.
settings:
  width: 1600
```

#### `height`

ウィンドウの高さ。

```yaml
# e.g.
settings:
  height: 800
```

## :book: サンプル

それ以外のサンプルについては [`examples/`](./examples/) ディレクトリを参照してください。

## :memo: ライセンス

[MIT License](./LICENSE)
