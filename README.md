# 競技プログラミング用ライブラリ

1. main.goの `func main()`の中に解法を書く。
2. [go-bundler](https://github.com/Atnuhs/go-bundler)を用いて一つのファイルのまとめる
3. 提出

# 注意事項
Goのバージョンで1.27.3を用いているためgo-bundlerを直接用いることができません。go-bundlerをcloneして
```bash
go get golang.org/x/tools/go/ast/astutil
go mod tidy
```
などを行うことでastutilのバージョンを最新にして対応させたのち
```bash
go install .
```
などをすれば扱えるようになると思います。
wslで使うことを想定して作成されているためbundle.shの中身は適宜自分の環境に合わせて用いてください
wslの場合はそのまま
```bash
./bundle.sh
```
すればbundleした上でコピペできるようになっています。
