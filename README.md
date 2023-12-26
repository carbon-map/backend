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
{"SQL_cmd":"SELECT amount, isPredict FROM carbonmap where year = ? and month = ? and city = ?","amount":"1155306600","isPredict":"false"}
```
