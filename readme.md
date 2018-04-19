# mqtt
1. 訂閱
2. 發送
3. 接收檔案、存檔


> 轉成 mp3, 64kbps 可將檔案壓縮較小

# how to use
1. `cd mqtt`
2. `docker-compose up -d`
3. 檢查 'app/audioSample/sender.go' 中要發送的檔案名稱
4. `docker exec audioPassByMqtt go run ex-audioPug.go`
5. 在 'app/audioSample' 中會發現新增 test.mp3
6. 透過 player 播放檔案