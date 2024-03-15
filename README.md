# SimpleTex-WebAPI
Transplant from github.com/scarletborder/simpleTex_Umi_Plugin

## Example
```python
import requests
file_path = "233.png"  # your image path
rec_mode_value = "auto"  # you can use formula/document/auto
with open(file_path, "rb") as fin:
    files = {"file": fin}
    data = {"rec_mode": rec_mode_value}
    resp = requests.post(
        """http://127.0.0.1:8080/upload""",
        files=files,
        data=data,
    )

print(resp.text)
```