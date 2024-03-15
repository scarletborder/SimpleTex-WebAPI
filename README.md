**本repo仅做学习用途**
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
```
'{"res":{"conf":0.934048593044281,"info":"C_m^n=C_{m-1}^{n-1}+C_{m-1}^n","rec_info":"org_no_bbox_easy","type":"formula"},"status":true}\n'
```

## Response Format
Response will be a json text. The text structure varies according to `rec_mode` you post.
All response in three `rec_modes` has one field

| field  | type | Comment                           |
| ------ | ---- | --------------------------------- |
| status | bool | Whether response status is normal |

If remote fails, resp will have additional field `errorMsg`

| field    | type   | Comment         |
| -------- | ------ | --------------- |
| errorMsg | string | detail of error |

When successfully post,  
general response will have these fields:
| field  | type   | comment                       |
| ------ | ------ | ----------------------------- |
| res    | struct | result of your uploaded image |
| status | bool   | true                          |

general `res` struct will have these arguments:
| field | type   | comment                                  |
| ----- | ------ | ---------------------------------------- |
| type  | string | recognition mode(`"formula"` or `"doc"`) |
### auto
When you set `rec_mode` as `auto`, the response will vary according to its `res/type` and will be the two formats below.

### formula
If your image is recognized as `formula` automatically
`res` struct
| field    | type   | comment       |
| -------- | ------ | ------------- |
| conf     | float  | confidential  |
| info     | string | Tex Text      |
| rec_info | string | no difference |
| type     | string | formula       |

### doc
If your image is recognized as `document` automatically or
`res` struct
| field | type   | comment |
| ----- | ------ | ------- |
| info  | struct |         |
| type  | string | doc     |

`info` struct

| field          | type   | comment       |
| -------------- | ------ | ------------- |
| bbox_number    | int    |               |
| section_number | int    |               |
| markdown       | string | result of OCR |
