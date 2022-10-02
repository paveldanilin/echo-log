set cpath=%~dp0

echo ">> linux_amd64"
docker run --rm -v %cpath%:/usr/src/app -w /usr/src/app golang:1.19 go build -v -o ./build/echo-log_linux_amd64 ./cmd/echo-log
echo "<< linux_amd64"

echo ">> windows_amd64.exe"
docker run --rm -v %cpath%:/usr/src/app -w /usr/src/app -e GOOS=windows -e GOARCH=386 golang:1.19 go build -v -o ./build/echo-log_windows_amd64.exe ./cmd/echo-log
echo "<< windows_amd64.exe"