from news_scraping.transform import (
    NewsDoc,
    get_str_comp,
)
from bs4.element import Tag
from datetime import datetime
from dateutil import parser
import re
from typing import Union
# from langdetect import detect
import uuid

def get_sa_news_doc(news_item: Tag) -> NewsDoc:
    title = get_str_comp(news_item, 'title')
    news_doc = NewsDoc(
        guid=get_str_comp(news_item, 'guid'),
        url=get_str_comp(news_item, 'link'),
        site='seekingalpha',
        site_full='https://seekingalpha.com/market-news',
        site_section=get_site_section(news_item),
        headline=title,
        title=title,
        body='', # going to use an empty string for now.
        ticker=get_ticker_symbol(news_item),
        tickers=[],
        published_str=get_str_comp(news_item, 'pubDate'),
        published_date=get_pub_datatime(news_item),
        language='en' # detect(title)
    )
    # convert the url guid to a uuid string. (unique to seeking alpha)
    news_doc.guid = get_doc_hash(news_doc)
    return news_doc

def get_doc_hash(news_doc: NewsDoc) -> str:
    return str(uuid.uuid3(uuid.NAMESPACE_URL, news_doc.guid))

def get_pub_datatime(news_item: Tag) -> datetime:
    pub_str = get_str_comp(news_item, 'pubDate')
    pub_date = parser.parse(pub_str).replace(tzinfo=None)
    return pub_date

def get_ticker_symbol(news_item: Tag) -> str:
    try:
        return get_str_comp(news_item, 'category').upper()
    except:
        return ""

def get_site_section(news_item: Tag) -> str:
    url = get_str_comp(news_item, 'link')
    site_section = '/'.join(url.split('/')[:4])
    return site_section
    