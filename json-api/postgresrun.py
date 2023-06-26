import os 
import conf
name_of_container = conf.name_of_postgres_container
postgres_password = conf.postgres_password  
os.system(
    "docker run --name {} -e POSTGRES_PASSWORD={} -p {}:{} -d {}".format(
        name_of_container, postgres_password,
        conf.dpvport, conf.dppport,
        conf.name_of_postgres_image)
)