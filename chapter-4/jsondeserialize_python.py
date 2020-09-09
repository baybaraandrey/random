from pprint import pprint
import json

if __name__ == '__main__':
	raw_data = input()
	metrics = json.loads(raw_data)
	
	print(type(metrics))
	print("Metrics from:", metrics.pop('hostname'))
	pprint(metrics)

