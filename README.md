# 【gin + gorm】ツイッターアプリのようなもの

```
go run main.go
```

## 注意事項
動かす前に、.envファイルをルートディレクトリに作って以下のように記述してください。`""`はいらないです。

```
mytweet_DBMS=mysql
mytweet_USER=<ユーザー名>
mytweet_PASS=<パスワード>
mytweet_DBNAME=<DB名>
```