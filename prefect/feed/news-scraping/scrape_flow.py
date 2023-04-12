from news_scraping.pipeline import scrape_rss
from news_scraping.load import IngestClient
from prefect import flow

@flow
def scrape_news():
    ingest_client = IngestClient.from_default('app.env')
    scrape_rss(ingest_client)

if __name__ == '__main__':
    scrape_news()