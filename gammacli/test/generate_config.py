import json
import os
import socket

def get_config(ip):
    return {
        "Environments": {
            "local": {
                "VirtualChain": 42,
                "Endpoints": [
                    "http://" + ip + ":8080"
                ]
            },
        }
    }

if __name__ == "__main__":
    ip = socket.gethostbyname(socket.gethostname())
    print json.dumps(get_config(ip))