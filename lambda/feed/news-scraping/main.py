import sys
sys.path.append('/Users/jst2702/Desktop/repos/projects/simple_sedge/sdk')
from news_scraping.pipeline import scrape_rss
from news_scraping.load import IngestClient

def main():
    client = IngestClient.from_default()
    scrape_rss(client)

if __name__ == '__main__':
    main()