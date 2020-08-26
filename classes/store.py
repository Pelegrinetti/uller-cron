import os
import time
import json


class Store:
    def __init__(self, store_folder):
        self.store_folder = store_folder

        os.makedirs(store_folder, exist_ok=True)
        self.file_path = os.path.join(store_folder, "data.json")

        if not os.path.exists(self.file_path):
            print('> Creating data file...')
            with open(self.file_path, 'w') as file:
                json.dump({'metrics': []}, file)
                file.close()
            print('> Data file created!')

    def store(self, values):
        with open(self.file_path, 'r+') as file:
            data = json.loads(file.read())
            data['metrics'].append(values)
            file.seek(0)
            file.write(json.dumps(data))
            file.truncate()
            file.close()

    def get(self):
        with open(self.file_path, 'r') as file:
            data = json.loads(file.read())
            return data['metrics']

    def truncate(self):
        with open(self.file_path, 'w') as file:
            json.dump({'metrics': []}, file)
            file.close()
