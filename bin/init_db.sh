# !/bin/bash
# docker-compose up -d
# cd docker || exit
docker-compose up -d && docker exec -i db_sso_demo psql -U postgres < ./docker/db_create_postgre.sql
# docker-compose up -d && docker exec -i demo_go_db mysql -uroot -ppassword demo_go < db_create.sql

# docker-compose up -d && docker exec -u postgres postgres_container psql demo_go postgres -f db_create.sql
# cat db_create_postgre.sql | docker exec -i db_sso_demo psql -U postgres