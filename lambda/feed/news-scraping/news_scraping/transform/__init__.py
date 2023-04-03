from news_scraping.common import (
    NewsDoc,
    rss_feeds_dict
)
from news_scraping.extract import get_str_comp

from .seeking_alpha import get_sa_news_doc
from .accesswire import get_aw_news_doc
from .benzinga import get_bz_news_doc
from .business_wire import get_bw_news_doc
from .globe_newswire import get_gn_news_doc
from .motley_fool import get_mf_news_doc
from .pr_newswire import get_pn_news_doc