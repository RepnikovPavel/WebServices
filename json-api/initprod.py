import os
import conf

if not os.path.exists(conf.executable_folder):
    os.makedirs(conf.executable_folder)

if not os.path.exists(conf.sqlite_folder):
    os.makedirs(conf.sqlite_folder)
