docker network create --driver bridge --subnet=172.16.12.0/16  allserver-net
docker network ls
docker network inspect allserver-net
::docker build -f Dockerfile_all  -t all . && docker run -d -p 0.0.0.0:3100:3100 -p 0.0.0.0:8081:8081  --name allserver  --network=allserver-net --ip 172.16.12.12 all
::docker run --name allserver all -d -p 31500:31500 31599:31599 3000:3000 3100:3100 31200:31200 31201:31201 31300:31300 31301:31301 31400:31400 31600:31600 31700:31700 31800:31800 8081:8081
docker build -f Dockerfile_all  -t all . && docker run -d -p 31500:31500 31599:31599 3000:3000 3100:3100 31200:31200 31201:31201 31300:31300 31301:31301 31400:31400 31600:31600 31700:31700 31800:31800 8081:8081 --name allserver all
::docker run -d -p 0.0.0.0:3100:3100 -p 0.0.0.0:8081:8081 --name allserver all
docker run -d -P --name allserver all
pause