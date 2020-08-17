from classes.store import Store
from datetime import datetime
import time
import requests
import json
import Adafruit_DHT
import settings

DHT_SENSOR = Adafruit_DHT.DHT22
DHT_PIN = 20


def init():
    print('> Running Uller cron!')
    store = Store('/tmp/uller/store')
    while True:
        humidity, temperature = Adafruit_DHT.read_retry(DHT_SENSOR, DHT_PIN)
        payload = {'temperature': temperature, 'humidity': humidity,
                   'created_at': datetime.now().strftime("%Y-%m-%d %H:%M:%S")}
        headers = {'content-type': 'application/json'}

        request = requests.post(settings.ULLER_API, data=json.dumps(payload), headers=headers)
        if request.status_code != 201:
            print('> Data stored because had an API error.')
            store.store(payload)

        time.sleep(2)


init()
