# backebnd
power by golang echo

# api
## "/" (GET): ???
## "/data" (GET): 取得資料庫中的資料
request example:
```
http://localhost:1323/data?year=2020&month=2&city=新北市
```
response example:
```
{"SQL_cmd":"SELECT amount FROM carbonmap where year = 2020 and month = 2 and city = '新北市'","amount":"1155306600"}
```
