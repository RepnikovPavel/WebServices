import os 
import conf
os.system(
    "kill -9 `pidof {}`".format(conf.exec_name)
)