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
$ make init
```
- .env.sampleの値を.envにコピー
- .envを書き換える

## 使う予定の外部ライブラリ
- chi
- [sqlc](https://github.com/kyleconroy/sqlc) or gorm
    - 渋川さんがお気に入りと仰られて居たので使痛い
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [github.com/caarlos0/env](https://github.com/caarlos0/env)
- air
- testify

