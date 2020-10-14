import pathlib
import requests
import file_util
import json
import time
import shutil
import json

config = json.loads(file_util.read_all_text('config.json'))

debug = config['debug']

if debug:
    host = 'http://localhost:80/api/DataLog?pageSize=99999'
else:
    host = 'https://smartlab.backend.117503445.top/api/DataLog?pageSize=99999'


dir_data = pathlib.Path('data')


def main():
    print(host)

    if dir_data.exists():
        shutil.rmtree(dir_data)

    response = requests.get(host)
    response.encoding = 'utf-8'
    text = response.text
    js = json.loads(text)

    for c in js['content']:
        text = c['content']

        id = c['id']
        page = c['page']
        file_name = f'{id}-{page}.csv'

        time_local = time.localtime(int(c['createdTimeStamp'])/1000)
        dt = time.strftime("%Y-%m-%d %H:%M:%S", time_local)

        openid = c['openid']

        text = f'Time,{dt}\n' + text
        text = f'Openid,{openid}\n' + text
        text = text.replace(',', '\t').replace(';', '')
        print(file_name)
        file_util.write_all_text(dir_data / file_name, text)


if __name__ == '__main__':
    print('每次运行脚本都会删除 data 文件夹的所有内容!!!!!')
    main()
