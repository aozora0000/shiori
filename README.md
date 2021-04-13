# Shiori(栞)

[![](https://img.youtube.com/vi/QOeibJtv1vk/0.jpg)](https://www.youtube.com/watch?v=QOeibJtv1vk)

前回までの表示位置を記憶し、次回起動時は続きからcatをしてくれるTerminalアプリです。
表示位置は$HOME/.config/shiori以下に保存されます。

## コマンド
### get,{default}
前回の表示位置から一度だけcatをします。
```terminal
$ shiori get example/test.txt
$ shiori example/test.txt
```

### watch
前回の表示位置からファイルの変更を検知してcatし続けます。
ファイル初期化があった場合は先頭から表示しなおします。
```terminal
$ shiori watch example/test.txt
```

### clear
表示位置の記憶ファイルを初期化します。
特定ファイルのみ初期化したい場合は、ファイルのURIを指定します。
```terminal
$ shiori clear #全削除
$ shiori clear example/test.txt #example/test.txtの表示位置のみ削除
```