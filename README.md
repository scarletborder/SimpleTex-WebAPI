# SimpleTex-WebAPI
Transplant from github.com/scarletborder/simpleTex_Umi_Plugin

## Example
```python

with open(file_path, "rb") as fin:
    files = {"file": fin}
    data = {"rec_mode": rec_mode_value}
    resp = requests.post(
        """http://127.0.0.1:8080/upload""",
        files=files,
        data=data,
        headers=headers,
        # proxies={
        #     "http": "http://10.10.1.10:3128",
        #     "https": "http://10.10.1.10:1080",
        # },
    )

```