# backebnd
power by golang echo

# api
## "/" (GET): ???
## "/data" (GET): 取得資料庫中的資料
request example:
```
http://{your backend url}/data?year=2020&month=2&city=%E6%96%B0%E5%8C%97%E5%B8%82
```
response example:
```
{"SQL_cmd":"SELECT amount, isPredict FROM carbonmap where year = 2020 and month = 2 and city = 新北市","amount":["1155306600"],"isPredict":["false"]}
```

request example:
```
http://{your backend url}/data?year=2020&month=all&city=%E6%96%B0%E5%8C%97%E5%B8%82
```
response example:
```
{"SQL_cmd":"SELECT amount, isPredict FROM carbonmap where year = 2020 and city = 新北市","amount":["1155306600","1435877378","1606787444","1232034707","1332323908","1570204669","1136375232","1222331287","1580172429","1167297867","1353532241","1355055406"],"isPredict":["false","false","false","false","false","false","false","false","false","false","false","false"]}
```
