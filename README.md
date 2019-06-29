# golang-web-base

Goプロジェクトのテンプレート

## 概要
- Gin + MySQL
- Docker環境
- 自動リロード

## プロジェクトの設置

`$GOPATH/src`内にプロジェクトが設置されていること
    - IntelliJの恩恵を受けるために

### GOPATH内にシンボリックリンクを作成する
```
ln -s `pwd` $GOPATH/src/github.com/$GIT_AUTHOR/
```

## 実行
```
docker-compose up api
```
