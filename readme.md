それぞれの動作確認

r.GET("/recruitments", getRecruitments)        // 募集要項一覧を表示
````
curl localhost:8080/recruitments
````

r.GET("/recruitments/:id", getRecruitmentById) // 指定した募集要項を表示
````
curl http://localhost:8080/recruitment/1 \
--header "Content-Type: application/json" \
--request "GET" \
--data '{"id": 1}'
````
or
````
curl http://localhost:8080/recruitment/1 \
--request "GET" \
--data '{"id": 1}'
````

r.POST("/recruitment", postRecruitment) // 募集要項を投稿
````
curl http://localhost:8080/recruitments \
--request "POST" \
--data '{"id": 4,"title": "タイトルD","need": "意欲が高い"}'
````
or
````
curl http://localhost:8080/recruitments \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"id": 4,"title": "タイトルD","need": "意欲が高い"}'
````

r.PATCH("/recruitment/:id", patchRecruitment)   // 募集要項を更新
````
curl http://localhost:8080/recruitment/1 \
--header "Content-Type: application/json" \
--request "PATCH" \
--data '{"title": "XXX","need": "YYY"}'
````

r.DELETE("/recruitment/:id", deleteRecruitment) // 募集要項を削除
````
curl http://localhost:8080/todos/2 \
--request "DELETE"
````
