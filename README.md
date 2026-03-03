# nvim-on-wezterm

WezTerm 上で Neovim を起動するための Windows 用の簡易実行ファイルです。
ファイルへの関連付けや、他ツールの外部エディタとして登録することを想定しています。

## 動作

```
nvim-on-wezterm.exe [path]
```

| 引数 | 動作 |
|------|------|
| なし | カレントディレクトリを cwd として WezTerm + Neovim を起動 |
| ファイルパス | そのファイルを Neovim で開き、cwd はファイルの親ディレクトリ |
| ディレクトリパス | Neovim の引数に渡し、cwd はそのディレクトリ |

内部では以下のコマンドを実行します。

```
wezterm-gui.exe start --cwd <cwd> -- nvim [path]
```

## 前提条件

- [WezTerm](https://wezfurlong.org/wezterm/) がインストールされ、`wezterm-gui.exe` が PATH に通っていること
- [Neovim](https://neovim.io/) がインストールされ、`nvim` が PATH に通っていること

## ビルド

```
go build -ldflags="-s -w -H=windowsgui" -o nvim-on-wezterm.exe
```

| フラグ | 説明 |
|--------|------|
| `-s -w` | デバッグ情報を除去してバイナリを小さくする |
| `-H=windowsgui` | 起動時にコンソールウィンドウを表示しない |

## 使用例

### ファイルエクスプローラーの「プログラムから開く」に登録

1. 任意のファイルを右クリックし、「プログラムから開く」→「別のプログラムを選択」を選択
2. `nvim-on-wezterm.exe` を指定

### コマンドラインから呼び出す

```powershell
# ファイルを開く
nvim-on-wezterm.exe C:\path\to\file.txt

# ディレクトリを開く
nvim-on-wezterm.exe C:\path\to\dir
```
