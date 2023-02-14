from news_modules import (
    NewsDoc,
    Site,
    get_str_comp,
    get_soup,
    rss_feeds_dict
)
from news_modules import Site
from bs4 import BeautifulSoup
from bs4.element import Tag
from datetime import datetime
from dateutil import parser
import re
from langdetect import detect
from typing import Iterable, List, Union


def get_pn_news_docs() -> Iterable[NewsDoc]:
    soup = get_soup(rss_feeds_dict[Site.PR_NEWSWIRE])
    news_docs = get_soup_docs(soup)
    return news_docs

def get_soup_docs(soup: BeautifulSoup) -> Iterable[NewsDoc]:
    items: List[Tag] = soup.find_all('item')
    for item in items:
        yield get_news_doc(item)

def get_news_doc(news_item: Tag) -> NewsDoc:
    title = get_str_comp(news_item, 'title')
    news_doc = NewsDoc(
        guid=get_str_comp(news_item, 'guid'),
        url=get_str_comp(news_item, 'link'),
        site='prnewswire',
        site_full='https://www.prnewswire.com/news-releases/news-releases-list/',
        site_section=get_site_section(news_item),
        headline=title,
        title=title,
        body='', # going to use an empty string for now.
        ticker=get_ticker_symbol(news_item),
        published_str=get_str_comp(news_item, 'pubDate'),
        published_date=get_pub_datatime(news_item),
        language=detect(title)
    )
    return news_doc

def get_pub_datatime(news_item: Tag) -> datetime:
    pub_str = get_str_comp(news_item, 'pubDate')
    pub_date = parser.parse(pub_str).replace(tzinfo=None)
    return pub_date

def get_ticker_symbol(news_item: Tag) -> Union[str, None]:
    desc_text = get_str_comp(news_item, 'description').upper()
    ticker_pattern = r"\(([A-Z.]+):(| )([A-Z.]+)\)"
    result = re.search(ticker_pattern, desc_text)
    if result:
        #ie. '(TSXV:IVI)' sub -> ' TSXV IVI ' strip -> 'TSXV IVI' split() -> ['TSXV','IVI']
        ticker = re.sub('[^A-Z.]+', ' ', result.group()).strip().split()[1]
        return ticker
    else: return None

def get_site_section(news_item: Tag) -> str:
    url = get_str_comp(news_item, 'link')
    site_section = '/'.join(url.split('/')[:4])
    return site_section
    