from news_scraping.load import IngestClient
from news_scraping.pipeline.get_docs import get_docs

def scrape_rss(client: IngestClient):
    for doc in get_docs():
        client.ingest_doc(doc)