# All except seeking alpha's were provided from rss.app
from enum import Enum


class Site(Enum):
    ACCESSWIRE = "accesswire"
    SEEKING_ALPHA = "seeking_alpha"
    BENZINGA = "benzinga"
    BUSINESS_WIRE = "business_wire"
    GLOBE_NEWSWIRE = "globe_newswire"
    MOTLEY_FOOL = "motley_fool"
    PR_NEWSWIRE = "pr_newswire"


rss_feeds_dict = {
    Site.ACCESSWIRE: "https://rss.app/feeds/QSnWPvwpfczVs1Pk.xml",
    Site.SEEKING_ALPHA: "https://seekingalpha.com/market_currents.xml",
    Site.BENZINGA: "https://rss.app/feeds/cofHV1JCTxNXtCYa.xml",
    Site.BUSINESS_WIRE: "https://rss.app/feeds/P3Byeak96AHjPjov.xml",
    Site.GLOBE_NEWSWIRE: "https://rss.app/feeds/DOHWgskBLxjeDlHc.xml",
    Site.MOTLEY_FOOL: "https://rss.app/feeds/cE62cwRkUjXRs745.xml",
    Site.PR_NEWSWIRE: "https://rss.app/feeds/uncRSaTuKbQntEhc.xml",
}