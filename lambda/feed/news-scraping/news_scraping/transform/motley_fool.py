from news_scraping.transform import (
    NewsDoc,
    get_str_comp,
)
from bs4.element import Tag
from datetime import datetime
from dateutil import parser
import re
from typing import Union, List
# from langdetect import detect


def get_mf_news_doc(news_item: Tag) -> NewsDoc:
    title = get_str_comp(news_item, 'title')
    tickers = get_ticker_symbol(news_item)
    news_doc = NewsDoc(
        guid=get_str_comp(news_item, 'guid'),
        url=get_str_comp(news_item, 'link'),
        site='fool',
        site_full='https://www.fool.com/investing-news/',
        site_section=get_site_section(news_item),
        headline=title,
        title=title,
        body='', # going to use an empty string for now.
        ticker=tickers[0],
        tickers=tickers,
        published_str=get_str_comp(news_item, 'pubDate'),
        published_date=get_pub_datatime(news_item),
        language='en' # detect(title)
    )
    return news_doc

def get_pub_datatime(news_item: Tag) -> datetime:
    pub_str = get_str_comp(news_item, 'pubDate')
    pub_date = parser.parse(pub_str).replace(tzinfo=None)
    return pub_date

def get_ticker_symbol(news_item: Tag) -> List[str]:
    desc_text = get_str_comp(news_item, 'title').upper()
    ticker_pattern = r"\$[A-Z]+"
    result = re.findall(ticker_pattern, desc_text)
    if result:
        tickers = [re.sub('[^A-Z.]+', '', t) for t in result]
        return tickers
    else: return []

def get_site_section(news_item: Tag) -> str:
    url = get_str_comp(news_item, 'link')
    site_section = '/'.join(url.split('/')[:4])
    return site_section
    