from selenium import webdriver
from selenium.webdriver.chrome.webdriver import WebDriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from webdriver_manager.chrome import ChromeDriverManager

service = Service(ChromeDriverManager().install())

def get_driver(headless=False) -> WebDriver:
    options = Options()
    options.headless = headless
    service = Service(ChromeDriverManager().install())

    driver = webdriver.Chrome(options=options, service=service)
    return driver