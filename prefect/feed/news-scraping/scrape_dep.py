from scrape_flow import scrape_news
from prefect.deployments import Deployment
from prefect.server.schemas.schedules import IntervalSchedule
from datetime import timedelta

schedule = IntervalSchedule(interval=timedelta(minutes=10))

deployment = Deployment.build_from_flow(
    flow=scrape_news,
    name="scrape-deployment",
    schedule=schedule
)

if __name__ == "__main__":
    deployment.apply() # type: ignore