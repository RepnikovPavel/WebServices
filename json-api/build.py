import os 
import conf
if conf.exec_name == 'main':
    raise Exception("the process name cannot be like this: {}".format(conf.exec_name))
os.system("go build -o {}".format(
    os.path.join(conf.executable_folder,conf.exec_name)
))
