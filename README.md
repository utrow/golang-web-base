# golang-web-base

Goプロジェクトのテンプレート

## 概要
- Gin + MySQL
  - ⚠ 最終的な構成イメージなので、レポジトリは未実装
- Docker環境
- 自動リロード

## 環境構築
まず、環境変数の設定を `.envrc.local` を参考に行う。
(direnv推奨)

### プロジェクトの設置

`$GOPATH/src`内にプロジェクトが設置されていること
    - IntelliJの恩恵を受けるために

#### GOPATH内にシンボリックリンクを作成する
```
ln -s `pwd` $GOPATH/src/github.com/$GIT_AUTHOR/
```

## 実行
```
docker-compose up
```
