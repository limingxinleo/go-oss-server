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