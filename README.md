# IIJmio プラン組み合わせ計算機

IIJmio の料金プランを複数回線で組み合わせ、指定した合計容量の範囲内で最安値を計算するツールです。

## 機能

- 回線数・最低/最高容量を指定して最安の組み合わせを表示
- 容量を 1GB 単位で一覧表示するモード
- CLI（Go バイナリ）と Web ブラウザ（WebAssembly）の両方に対応

## 料金プラン

| 容量 | 月額 |
|------|------|
| 2GB  | ¥850 |
| 5GB  | ¥950 |
| 10GB | ¥1,400 |
| 15GB | ¥1,600 |
| 25GB | ¥2,000 |
| 35GB | ¥2,400 |
| 45GB | ¥3,300 |
| 55GB | ¥3,900 |

> 複数回線割引: ¥100/回線/月

## 使い方

### CLI

```sh
# ビルド
go build -o iij-kumiawase .

# 例: 4回線で合計 25〜45GB の最安値を計算
./iij-kumiawase -lines 4 -min 25 -max 45

# 1GB 単位で一覧表示
./iij-kumiawase -lines 4 -min 25 -max 45 -all
```

#### オプション

| フラグ | デフォルト | 説明 |
|--------|-----------|------|
| `-lines` | 4 | 回線数 |
| `-min` | 25 | 最低合計容量 (GB) |
| `-max` | 45 | 最高合計容量 (GB) |
| `-all` | false | min〜max を 1GB 単位で一覧表示 |

### Web (WebAssembly)

`index.html`・`main.wasm`・`wasm_exec.js` を同じディレクトリに置き、HTTP サーバーで配信します。

```sh
python3 -m http.server 8080
```

ブラウザで `http://localhost:8080` を開いてください。

#### WASM のビルド

```sh
GOOS=js GOARCH=wasm go build -o main.wasm .
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```
