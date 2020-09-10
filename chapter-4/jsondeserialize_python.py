from pprint import pprint
import json

if __name__ == '__main__':
	while True:
		raw_data = input()
		metrics = json.loads(raw_data)
		print("Metrics from:", metrics.pop('hostname'))
		pprint(metrics)

