# backebnd
power by golang echo

# api
## 給定年、月、城市，查詢當年度某月某城市的碳排放量資料
### request
| HTTP method | request URL |
|-------------|-------------|
| GET         | `/data`     |

### request parameters
| 參數名稱   | 類型     | 必要性 | 說明         |
|-----------|----------|--------|--------------|
| year      | string   | 是     | 選取年分 |
| month     | string   | 是     | 選取月分 |
| city      | string   | 是     | 選取城市 |

### request example:
```
http://{your backend url}/data?year=2023&month=12&city=嘉義市
```

### response parameters
| 參數名稱     | 類型      | 必要性 | 說明                                |
|-------------|-----------|--------|-------------------------------------|
| SQL_cmd     | string    | 是     | 後端 query database 的指令           |
| amount      | string[]  | 是     | 碳排放量，只會有一筆資料              |
| isPredict   | string[]  | 是     | 該筆資料是否是預測的，只會有一筆資料  |

response example:
```
{
   "SQL_cmd":"SELECT amount, isPredict FROM carbonmap where year = 2023 and month = 12 and city = 嘉義市",
   "amount":[
      "97413139"
   ],
   "isPredict":[
      "true"
   ]
}
```

## 給定年、城市，查詢當年度某城市每個月的碳排放量資料
### request
| HTTP method  | request URL |
|-----------   |-------------|
| GET          | `/data`     |

### request parameters
| 參數名稱  | 類型     | 必要性 | 說明              |
|----------|----------|--------|-------------------|
| year     | string   | 是     | 選取年分          |
| month    | string   | 是     | 本參數固定為 all  |
| city     | string   | 是     | 選取城市          |

### request example:
```
http://{your backend url}/data?year=2023&month=all&city=嘉義市
```

### response parameters
| 參數名稱    | 類型      | 必要性 | 說明                        |
|------------|-----------|--------|-----------------------------|
| SQL_cmd    | string    | 是     | 後端 query database 的指令  |
| amount     | string[]  | 是     | 碳排放量                    |
| isPredict  | string[]  | 是     | 該筆資料是否是預測的         |

response example:
```
{
   "SQL_cmd":"SELECT amount, isPredict FROM carbonmap where year = 2023 and city = 嘉義市",
   "amount":[
      "92619574",
      "90109625",
      "97413139",
      "100340432",
      "88360140",
      "99573918",
      "87089017",
      "80552324",
      "77938161",
      "86893505",
      "101058559",
      "89348585"
   ],
   "isPredict":[
      "true",
      "true",
      "true",
      "false",
      "false",
      "false",
      "false",
      "false",
      "false",
      "false",
      "false",
      "false"
   ]
}
```
