# Go OSS Server

## Docker启动

```
docker run -p 8080:8080 \
-e END_POINT=your_end_point \
-e ACCESS_KEY_ID=your_key_id \
-e ACCESS_KEY_SECRET=your_key_secret \
limingxinleo/go-oss-server:latest
```

测试
```
curl -X POST http://localhost:8080/simple_handler/your_bucket\?object\=your_object\
  -F "file=@/Users/xxx/Downloads/picture/test.jpeg" \
  -H "Content-Type: multipart/form-data"
  
curl -X POST http://localhost:8080/simple_handler/your_bucket \
  -F "file=@/Users/xxx/Downloads/picture/test.jpeg" \
  -H "Content-Type: multipart/form-data"
```

## stack 样例
```yaml
version: '3.7'
services:
  go_app:
    image: $REGISTRY_URL/$PROJECT_NAME:test
    environment:
      - "APP_PROJECT=go_app"
      - "APP_ENV=test"
    ports:
      - 8080:8080
    deploy:
      replicas: 1
      restart_policy:
        condition: on-failure
        delay: 5s
        max_attempts: 5
      update_config:
        parallelism: 2
        delay: 5s
        order: start-first
    networks:
      - go_app_net
    configs:
      - source: go_app_v1.0
        target: /.env
configs:
  go_app_v1.0:
    external: true
networks:
  go_app_net:
    external: true
```