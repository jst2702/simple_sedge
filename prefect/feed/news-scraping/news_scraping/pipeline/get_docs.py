from news_scraping.common import (
    NewsDoc,
    Site,
    rss_feeds_dict,
)
from news_scraping.transform import (
    get_aw_news_doc,
    get_bw_news_doc,
    get_bz_news_doc,
    get_gn_news_doc,
    get_mf_news_doc,
    get_pn_news_doc,
    get_sa_news_doc
)
from news_scraping.extract import get_soup
from bs4 import BeautifulSoup, Tag
from typing import Iterable, List, Callable

def get_docs() -> Iterable[NewsDoc]:
    yield from get_aw_news_docs()
    yield from get_bw_news_docs()
    yield from get_bz_news_docs()
    yield from get_gn_news_docs()
    yield from get_mf_news_docs()
    yield from get_pn_news_docs()
    yield from get_sa_news_docs()

def get_aw_news_docs() -> Iterable[NewsDoc]:
    soup = get_soup(rss_feeds_dict[Site.ACCESSWIRE])
    news_docs = _get_soup_docs(soup, get_aw_news_doc)
    return news_docs

def get_bw_news_docs() -> Iterable[NewsDoc]:
    soup = get_soup(rss_feeds_dict[Site.BUSINESS_WIRE])
    news_docs = _get_soup_docs(soup, get_bw_news_doc)
    return news_docs

def get_bz_news_docs() -> Iterable[NewsDoc]:
    soup = get_soup(rss_feeds_dict[Site.BENZINGA])
    news_docs = _get_soup_docs(soup, get_bz_news_doc)
    return news_docs

def get_gn_news_docs() -> Iterable[NewsDoc]:
    soup = get_soup(rss_feeds_dict[Site.GLOBE_NEWSWIRE])
    news_docs = _get_soup_docs(soup, get_gn_news_doc)
    return news_docs

def get_mf_news_docs() -> Iterable[NewsDoc]:
    soup = get_soup(rss_feeds_dict[Site.MOTLEY_FOOL])
    news_docs = _get_soup_docs(soup, get_mf_news_doc)
    return news_docs

def get_pn_news_docs() -> Iterable[NewsDoc]:
    soup = get_soup(rss_feeds_dict[Site.PR_NEWSWIRE])
    news_docs = _get_soup_docs(soup, get_pn_news_doc)
    return news_docs

def get_sa_news_docs() -> Iterable[NewsDoc]:
    soup = get_soup(rss_feeds_dict[Site.SEEKING_ALPHA])
    news_docs = _get_soup_docs(soup, get_sa_news_doc)
    return news_docs

def _get_soup_docs(
    soup: BeautifulSoup, 
    call_get_doc: Callable[[Tag], NewsDoc]) -> Iterable[NewsDoc]:

    items: List[Tag] = soup.find_all('item')
    for item in items:
        yield call_get_doc(item)