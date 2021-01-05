# README

### Requirement

- 미국 주식 시장의 데이터를 가져올 수 있어야 한다.
- 지난 6개월 간의 데이터 중 가장 큰 이익을 얻을 수 있는 매수 날짜와 매도 날짜를 찾아서 반환 할 수 있어야 한다.


#### User story

- User는 원하는 stock symbol code를 입력하면 데이터를 제공 받을 수 있다.

### 설계

#### Flow Chart

![flowchart.png](https://github.com/Kyejoon-Lee/OOP_Python/blob/master/Baseball/flowchart.png?raw=true)

#### 개발 도구 및 외부 API

- [Golang echo](https://echo.labstack.com/) 프레임워크를 사용하여 개발한다.
- https://github.com/markcheno/go-quote 의 API를 사용, Yahoo finance를 통해 주식 data 요청.

```
go get github.com/labstack/echo/middleware

go get github.com/markcheno/go-quote

go install main

go install front
```

- 정적 페이지를 분리해 놓았기 때문에 main, front를 각각 실행주어야한다.

```
bin/main

bin/front
```



#### 알고리즘

- 일별 최고가, 최저가를 고려한 알고리즘
  - 미국의 경우 상한가, 하한가가 존재 하지 않기 때문에 하루만에 최대의 이익을 얻을 수도 있다. 그러므로 종가를 기준으로 하지 않고 low 값을 지속적으로 갱신 후 비교를 진행 한 후 결과를 도출한다.
  - 일별 최대 하한 가격의 시간과 상한 가격의 시간이 제공 되지 않다는 점이 존재한다.
- 종가를 기준으로 한 알고리즘

#### Design

- AWS를 사용하여 간단한 웹페이지와 결과값을 창으로 나타낼 수 있도록한다.

- web 주소(접속 가능) : http://52.78.172.184:8080/

- Go
  - 총 4가지로 구성이 되어있다.
- controller
    - webpagecontroller
     - 간단하게 front와 back-end를 구분 짓기 위해 구성하였다.
       - 구성은 http://52.78.172.184:8081/index.html 로 redirect 하도록 설계하였다.
  - model
    - myinformation
       - web server로 반환을 위한 구조체로써 사용되어진다.
       - Buy ,Sell, Profit, Symbol로써 구성되어진다.
    - reqdata
      - ReqData(string) : user 요청으로 symbol을 받아 Yahoo Finance로 부터 6개월간의 주식 기록을 받아온다.
  - service
    - finddate
      - FindDate(quote.Quote) : 6개월간의 주식 기록을 일별 최고가, 최저가를 바탕으로 결과값 도출.
      - FindDateClose(quote.Quote) : 6개월간의 주식 기록을 종가를 바탕으로 결과값 도출.
  - api
    - handlers
      - symbolhandlers
        - GetData(echo.Context) : user로부터 symbol code를 받아와 결과를 return 해준다. 

#### API

- [getData](https://documenter.getpostman.com/view/13164055/TVYAi1yq)

#### 실행 방법

- 외부 라이브러리를 import 한 후 install 한다.

```
go get github.com/labstack/echo/middleware

go get github.com/markcheno/go-quote

go install main

go install front
```

- 정적 페이지를 분리해 놓았기 때문에 main, front를 각각 실행주어야한다.

```
bin/main

bin/front
```



#### Test

- 간단한 test를 실행한 결과 값을 첨부합니다.

![testcase.png](https://github.com/Kyejoon-Lee/OOP_Python/blob/master/Baseball/testcase.png?raw=true)

- 총 8개의 Symbol code가 있다. 이 중 APPLE, SAMSUNG 을 제외하고는 test에 대해 통과를 해야한다.

![response.png](https://github.com/Kyejoon-Lee/OOP_Python/blob/master/Baseball/response.png?raw=true)

- 예상처럼 비정상적인 Symbol code에 대해서 404 error를 도출하는것을 볼 수 있으며 대소문자 구별없이 잘 작동하는 것을 볼 수 있다.





