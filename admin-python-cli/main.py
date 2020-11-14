import pathlib
import requests
import file_util
import json
import time
import shutil
import json

config = json.loads(file_util.read_all_text('config.json'))

debug = config['debug']


dir_data = pathlib.Path('data')

meta = [["获取所有 DataLog", "get_data_log"], ["发布公告", "send_bulletin"],["查看公告","get_bulletins"],["删除公告","delete_bulletin"]]

if debug:
    host = 'http://localhost:80/api'
else:
    host = 'https://smartlab.backend.117503445.top/api'


def get_data_log():
    print('每次运行脚本都会删除 data 文件夹的所有内容!!!!!')

    global host
    url = f"{host}/DataLog?pageSize=99999"
    print(url)

    if dir_data.exists():
        shutil.rmtree(dir_data)

    response = requests.get(url)
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


def login(session):
    global host
    global config
    url = f"{host}/user/login"
    user_login_in = {
        "Username": config["username"],
        "Password": config["password"]
    }

    res = session.post(url, json.dumps(user_login_in))
    if res.status_code != 200:
        print('login failed')
        print(res.status_code)
        print(res.text)
        exit()
    js = json.loads(res.text)
    session.headers["Authorization"] = "Bearer " + js['token']


def send_bulletin():

    s = requests.session()

    login(s)

    global host
    url = f"{host}/Bulletin"
    print(url)

    image_url =input( 'please input image url\n')
    title = input('please input title\n')
    bulletin_in = {
        "ImageUrl": image_url,
        "Title": title
    }

    s.post(url, json.dumps(bulletin_in))

def get_bulletins():
    s = requests.session()

    # login(s)
    global host
    url = f"{host}/Bulletin"
    print(url)
    r = s.get(url)
    print(r.text)



def delete_bulletin():
    get_bulletins()
    s = requests.session()

    login(s)
    global host
    id = input('please input id you want to delete\n')
    url = f"{host}/Bulletin/{id}"


    s.delete(url)


def main():
    global meta
    for i in range(len(meta)):
        print(i, meta[i][0])
    index = int(input('请输入序号 '))

    func_name = meta[index][1]

    eval(func_name+'()')


if __name__ == '__main__':
    main()
