docker build -t renlu/gmm .
docker stop gmm
docker rm gmm
docker run -d --name=gmm --restart=always -p 7654:7654 renlu/gmm