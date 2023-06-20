import os
import conf
if not os.path.exists(conf.path_to_downloads):
    os.makedirs(conf.path_to_downloads)
os.system(
    "wget {} -P {}".format(conf.csv_link,conf.path_to_downloads)
)