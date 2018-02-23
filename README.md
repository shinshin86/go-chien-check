# go-chien-check



## これはなに？ 

電車の遅延情報を[YAHOO!路線情報 - 運行情報](https://transit.yahoo.co.jp/traininfo/top)を用いてコマンドラインからチェックするためのCLIツール



## 使い方

※下記は暫定的に記載している。ここは後で整える(TODO)

```bash
go run main.go
```



### 調べたい地域の路線について

データ取得先である[YAHOO!路線情報 - 運行情報](https://transit.yahoo.co.jp/traininfo/top)の仕様により、対応する地域のエリア番号をURLに設定することにより、対象地域の遅延情報を調べることができます。

このエリア番号については、現在はソース内に直に設定するようになっています。

※サイトの仕様が変更されることにより、今後使用不可能になる可能性があります。



| エリア番号 | 対応地域 |
| ---------- | -------- |
| 2          | 北海道   |
| 3          | 東北     |
| 4          | 関東     |
| 5          | 中部     |
| 6          | 近畿     |
| 8          | 中国     |
| 9          | 四国     |
| 7          | 九州     |



## データ参照先

[YAHOO!路線情報 - 運行情報](https://transit.yahoo.co.jp/traininfo/top)



## 使用上の注意

[YAHOO!路線情報 - 運行情報](https://transit.yahoo.co.jp/traininfo/top)は特にWebスクレイピングは禁止されていないようですが、データ取得は必要最低限にとどめてください。<br>

なお、このスクリプトは一回のコマンドで一度だけデータ取得のためにアクセスをするようになっています。

参照 : [robots.txt](https://transit.yahoo.co.jp/robots.txt)

