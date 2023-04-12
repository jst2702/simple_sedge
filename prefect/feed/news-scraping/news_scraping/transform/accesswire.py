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

def get_aw_news_doc(news_item: Tag) -> NewsDoc:
    title = get_str_comp(news_item, 'title')
    news_doc = NewsDoc(
        guid=get_str_comp(news_item, 'guid'),
        url=get_str_comp(news_item, 'link'),
        site='accesswire',
        site_full='https://www.accesswire.com/newsroom',
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
    return news_doc

def get_pub_datatime(news_item: Tag) -> datetime:
    pub_str = get_str_comp(news_item, 'pubDate')
    pub_date = parser.parse(pub_str).replace(tzinfo=None)
    return pub_date

def get_ticker_symbol(news_item: Tag) -> str:
    desc_text = get_str_comp(news_item, 'description').upper()
    ticker_pattern = r"\(([A-Z.]+):(| )([A-Z.]+)\)"
    result = re.search(ticker_pattern, desc_text)
    if result:
        #ie. '(TSXV:IVI)' sub -> ' TSXV IVI ' strip -> 'TSXV IVI' split() -> ['TSXV','IVI']
        ticker = re.sub('[^A-Z.]+', ' ', result.group()).strip().split()[1]
        return ticker
    else: return ""

def get_site_section(news_item: Tag) -> str:
    url = get_str_comp(news_item, 'link')
    site_section = '/'.join(url.split('/')[:4])
    return site_section
    