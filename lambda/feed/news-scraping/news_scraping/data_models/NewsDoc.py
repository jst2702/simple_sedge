from dataclasses import dataclass
from datetime import datetime
from typing import List, Union

@dataclass
class NewsDoc:
    guid: str
    url: str
    site: str
    site_full: str
    site_section: str
    headline: str
    title: str
    body: str
    ticker: Union[None, str, List[str]]
    published_str: str
    published_date: datetime
    language: str

@classmethod
def gen_guid(cls, **kwargs) -> NewsDoc:
    kwargs['guid'] = kwargs['url']
    return cls(**kwargs)