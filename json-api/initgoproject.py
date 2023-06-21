import os

pwd = os.getcwd()
current_dir_name = pwd.split('/')[-1]
os.system("go mod init {} && go mod tidy".format(current_dir_name))
