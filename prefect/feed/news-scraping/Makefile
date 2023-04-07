# Makefile

login:
	prefect cloud login -k ${PREFECT_API_KEY}

# deploy:
# 	prefect deployment build -n scrape-deployment scrape_news.py:scrape_news
# 	prefect deployment apply scrape_news-deployment.yaml
# 	make agent

deploy:
	py39 scrape-dep.py

agent:
	prefect agent start -p 'default-agent-pool'