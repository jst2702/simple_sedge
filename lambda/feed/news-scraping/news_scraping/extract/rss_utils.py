import requests
from bs4 import BeautifulSoup
from bs4.element import Tag


def get_soup(feed_url: str) -> BeautifulSoup:
    """gets the soup for the news pieces currently on the rss feed.
    Returns:
        BeautifulSoup: _description_
    """
    html_doc = requests.get(feed_url).content
    soup = BeautifulSoup(html_doc, 'xml')
    return soup

def get_str_comp(tag: Tag, key: str) -> str:
    elem = tag.find(key)
    if isinstance(elem, Tag) and elem.string:
        return elem.string
    else:
        raise ValueError(f"could not find string component for: '{key}'")