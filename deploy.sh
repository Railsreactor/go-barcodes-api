go get
GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o main_linux
scp ./main_linux "${DEPLOY_USER}@${DEPLOY_HOST}:~/barcodes_api/main_linux_new"
ssh "${DEPLOY_USER}@${DEPLOY_HOST}" 'pkill main_linux'
ssh "${DEPLOY_USER}@${DEPLOY_HOST}" 'mv ~/barcodes_api/main_linux_new ~/barcodes_api/main_linux'

(ssh "${DEPLOY_USER}@${DEPLOY_HOST}" 'cd ~/barcodes_api/ && PORT=8081 GIN_MODE=release nohup ./main_linux > log.log') & sleep 5 ; kill $!
echo 'Done'
