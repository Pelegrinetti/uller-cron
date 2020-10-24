from classes.store import Store
from datetime import datetime
import time
import requests
import json
import Adafruit_DHT
import settings

DHT_SENSOR = Adafruit_DHT.DHT22
DHT_PIN = 20


class Metric:
    def __init__(self):
        self.store = Store('/tmp/uller/store')

    def getSensorData(self):
        humidity, temperature = Adafruit_DHT.read_retry(DHT_SENSOR, DHT_PIN)
        return {'temperature': temperature, 'humidity': humidity,
                'created_at': datetime.now().strftime("%Y-%m-%d %H:%M:%S")}

    def sendData(self, data,  many=False):
        headers = {'content-type': 'application/json'}
        request = requests.post(
            settings.ULLER_API + ('/metrics' if many else '/metric'), data=json.dumps(data), headers=headers, params={'key': settings.ULLER_API_KEY})
        return request

    def localStore(self, data):
        self.store.store(data)

    def getLocalData(self):
        return self.store.get()

    def truncateData(self):
        self.store.truncate()


def init():
    print('> Running Uller cron!')
    metric = Metric()
    while True:
        try:
            data = metric.getSensorData()
            request = metric.sendData(data, False)
            if request.status_code != 201:
                print('> Data stored because had an API error.')
                metric.localStore(data)
                pass

            if request.status_code == 201:
                print('> Sending metrics to API.')
                data = metric.getLocalData()
                if len(data) > 0:
                    request = metric.sendData(data, True)
                    if request.status_code == 201:
                        metric.truncateData()
                        print('> Local data has been sent to the API.')
                        pass
                    else:
                        print('> Error while trying to send local data to API.')
                        pass
                else:
                    print('> Nothing to send to Uller API.')
                    pass
                pass
            pass
        except requests.exceptions.ConnectionError:
            print('> Couldn\'t connect with Uller API.')
            metric.localStore(data)
            pass

        time.sleep(2)


init()
