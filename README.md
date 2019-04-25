# Go OSS Server

## Docker启动

docker run -p 8080:8080 \
-e END_POINT:your_end_point \
-e ACCESS_KEY_ID:your_key_id \
-e ACCESS_KEY_SECRET:your_key_secret \
limingxinleo/go-oss-server:latest