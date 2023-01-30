# 開発手順
- jiraにチケット作る
- feature/チケットに対応したbranch作成
- レビュワー指定してpull request
  - templateに沿って記入
- (一応)slackのtech-pj-account-githubにメンション
- okならmerge

# 開発環境 setup

```shell
$ touch .env
$ make build
$ make up
```
- .env.sampleの値を.envにコピー
- .envを書き換える

## 使う予定の外部ライブラリ
- chi
  - ルーター
  - 多分なんでもいい
- [sqlc](https://github.com/kyleconroy/sqlc) or gorm
    - sqlcは渋川さんがお気に入りと仰られて居たので使ってみたい,,,
- golang-migrate
  - マイグレーション
  - 多分なんでもいい
- [github.com/caarlos0/env](https://github.com/caarlos0/env)
  - 環境変数読み込み
- air
  - 変更を感知して対象path(今回はtmp)以下にbuildし直す
- testify

## フォルダ設計
```shell
```
- 出来そうな気がするので4層で作ってみる

## 環境変数
