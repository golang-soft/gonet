docker build -f Dockerfile_netgate  -t netgate . && docker run --name netgateserver netgate
pause
::docker run -d -p 0.0.0.0:3100:3100 -p 0.0.0.0:8081:8081 --name netgateserver netgate