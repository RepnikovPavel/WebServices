import os
executable_folder = '/home/user/WebServices/json-api-Exec'
exec_name = 'un1234'
base_db_folder = '/home/user/WebServices/json-api-DataBase'
sqlite_folder = os.path.join(base_db_folder,'SQLite')
path_to_downloads = os.path.join(base_db_folder,'Downloads')
csv_link = "https://www.fdic.gov/bank/individual/failed/banklist.csv"

# docker postgres database 
name_of_postgres_container  = "postgrescont"
name_of_postgres_image = "postgres"
postgres_password  = "postgres1234"
dpvport = "5432"
dppport = "5432"