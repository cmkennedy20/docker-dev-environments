# docker-dev-environments
* Project to illustrate the basics of how a docker development environment can be configured

# Important Commands 
| Command Name | Command | 
| :------------| :------------|
| Docker build command |docker build -f `dockerfile` -t repo_name:tag . |
| Docker execute command | docker exec -it $(docker ps | cut -d " " -f 1) bash |
| Docker compose up | docker compose -f compose.yml up -d |
| Docker compose down | docker compose -f compose.yml down |  
