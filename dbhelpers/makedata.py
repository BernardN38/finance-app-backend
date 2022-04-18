import math
from models import Transaction, User
from random import randrange, uniform
from sqlalchemy.orm import Session
from sqlalchemy import create_engine
from pprint import pprint


engine = create_engine("postgresql:///finance_app", echo=False, future=True)


TYPES = ["balance", "deposit", "withdrawal", "description"]

MONTHS = {
    "Jan": [],
    "Feb": [],
    "Mar": [],
    "Apr": [],
    "May": [],
    "Jun": [],
    "Jul": [],
    "Aug": [],
    "Sep": [],
    "Oct": [],
    "Nov": [],
    "Dec": [],
}
DAYS_IN_MONTH = {
    "Jan": 31,
    "Feb": 28,
    "Mar": 31,
    "Apr": 30,
    "May": 31,
    "Jun": 30,
    "Jul": 31,
    "Aug": 31,
    "Sep": 30,
    "Oct": 31,
    "Nov": 30,
    "Dec": 31,
}
MONTH_INDEX = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
]


balance = 1000
investment_total = 5000
intrest_earned = 0
DAYS_IN_YEAR = 365


def year_transactions(
    day=1,
    month=0,
    year="2021",
    salary=80000,
    tax_rate=22,
    investment_percent=7,
    retirement_percent=5,
    daily_spending_limit=60.00,
    mortgage_payment=1150,
    car_payment=350,
    utilities=150,
):

    current_day = day
    payday_interval = 14
    days_until_payday = 0

    # valid date check
    while current_day < 365:
        # valid date check
        if current_day > DAYS_IN_MONTH[MONTH_INDEX[month]]:
            current_day -= DAYS_IN_MONTH[MONTH_INDEX[month]]
            month += 1
        if current_day > 365 or month > 11:
            break

        # create date string
        date = f'{month+1}-{current_day}-{year}'

        # payday
        if days_until_payday == 0:
            pay_amount = pay_day(date, month, salary, tax_rate)

            make_retirement_contributions(
                date, month, pay_amount, retirement_percent)
            make_investment(date, month, pay_amount, investment_percent)

            days_until_payday += payday_interval

        # pay bills at start of month
        if current_day == 1:
            calculate_intrest_on_investment(date, month, randrange(-5, 6))
            pay_bills(date, month, mortgage_payment, car_payment, utilities)

        # make daily purchases & withdrawals
        make_purchases(date, month, daily_spending_limit)

        # calculate emergency
        check_emergency(date, month)

        # change state to next day
        days_until_payday -= 1
        current_day += 1

    print(intrest_earned)


def init():
    MONTHS = {
        "Jan": [],
        "Feb": [],
        "Mar": [],
        "Apr": [],
        "May": [],
        "Jun": [],
        "Jul": [],
        "Aug": [],
        "Sep": [],
        "Oct": [],
        "Nov": [],
        "Dec": [], }
    balance = 1000
    investment_total = 500
    intrest_earned = 0


def pay_day(date, month, salary, tax_rate):
    global balance
    tax_amount = salary // 27 * (tax_rate/100)
    pay_amount = round(salary/27, 2)

    balance += pay_amount
    MONTHS[MONTH_INDEX[month]].append(
        {
            "date": date,
            "balance":  round(balance, 2),
            "deposit": pay_amount,
            "withdrawal": 0,
            "description": "pay_day",
            "investment_total":  round(investment_total, 2), }
    )
    balance -= tax_amount
    MONTHS[MONTH_INDEX[month]].append(
        {
            "date": date,
            "balance":  round(balance, 2),
            "deposit": 0,
            "withdrawal": tax_amount,
            "description": "taxes",
            "investment_total": round(investment_total, 2), }
    )
    return pay_amount


def pay_bills(date, month, mortgage_payment, car_payment, utilities):
    global balance
    balance -= car_payment
    MONTHS[MONTH_INDEX[month]].append(
        {
            "date": date,
            "balance":  round(balance, 2),
            "deposit": 0,
            "withdrawal": car_payment,
            "description": "car_payment",
            "investment_total": round(investment_total, 2), }
    )
    balance -= mortgage_payment
    MONTHS[MONTH_INDEX[month]].append(
        {
            "date": date,
            "balance":  round(balance, 2),
            "deposit": 0,
            "withdrawal": mortgage_payment,
            "description": "mortgage_payment",
            "investment_total": round(investment_total, 2), }
    )

    balance -= utilities
    MONTHS[MONTH_INDEX[month]].append(
        {
            "date": date,
            "balance": round(balance, 2),
            "deposit": 0,
            "withdrawal": utilities,
            "description": "utilities",
            "investment_total": round(investment_total, 2), }
    )


def calculate_intrest_on_investment(date, month, market_variance):
    global investment_total, intrest_earned
    intrest = round((investment_total * (market_variance / 100)),2)
    investment_total += intrest
    intrest_earned += intrest
    MONTHS[MONTH_INDEX[month]].append(
        {
            "date": date,
            "balance":  round(balance, 2),
            "deposit": 0,
            "withdrawal": 0,
            "description": "investment_total",
            "investment_total": round(investment_total, 2), }
    )
    MONTHS[MONTH_INDEX[month]].append(
        {
            "date": date,
            "balance":  round(balance, 2),
            "deposit": intrest,
            "withdrawal": 0,
            "description": "investment_intrest",
            "investment_total": round(investment_total, 2), }
    )


def make_investment(date, month, pay_amount, investment_percent):
    investment_amount = pay_amount * (investment_percent / 100)
    global investment_total, balance
    investment_total += investment_amount
    if balance > investment_amount:
        balance -= investment_amount
        MONTHS[MONTH_INDEX[month]].append(
            {
                "date": date,
                "balance":  round(balance, 2),
                "deposit": 0,
                "withdrawal": round(investment_amount, 2),
                "description": "investment",
                "investment_total": round(investment_total, 2), }
        )


def make_purchases(date, month, daily_spending_limit):
    global balance
    day_spent_amount = 0
    while daily_spending_limit > day_spent_amount:
        purchase_amount = round(uniform(1, daily_spending_limit),2)
        if balance*1.50 > purchase_amount:
            balance -= purchase_amount
            day_spent_amount += purchase_amount
            MONTHS[MONTH_INDEX[month]].append(
                {
                    "date": date,
                    "balance":  round(balance, 2),
                    "deposit": 0,
                    "withdrawal": purchase_amount,
                    "description": "purchase",
                    "investment_total": round(investment_total, 2), }
            )
        else:
            break


def check_emergency(date, month):
    global balance
    emergency = False
    if randrange(1, 101) > 99:
        emergency = True

    if emergency == True:
        emergency_amount = randrange(200, 1750)
        print("emergency amount--", emergency_amount)
        balance -= emergency_amount
        MONTHS[MONTH_INDEX[month]].append(
            {
                "date": date,
                "balance":  round(balance, 2),
                "deposit": 0,
                "withdrawal": emergency_amount,
                "description": "emergency",
                "investment_total": round(investment_total, 2), }
        )


def make_retirement_contributions(date, month, pay_amount, retirement_percent):
    retirement_contribution = pay_amount * retirement_percent / 100
    global balance
    global investment_total
    balance -= retirement_contribution
    investment_total += retirement_contribution
    MONTHS[MONTH_INDEX[month]].append(
        {
            "date": date,
            "balance": round(balance, 2),
            "deposit": 0,
            "withdrawal": retirement_contribution,
            "description": "retirement_contribution",
            "investment_total": round(investment_total, 2), }
    )


def add_data_db():
    year_transactions()
    with Session(engine) as session:
        transactions = []

        # Transaction.__table__.drop(engine)
        # User.__table__.drop(engine)
        # User.__table__.create(engine)
        # Transaction.__table__.create(engine)
        for month in MONTHS:
            for transaction in MONTHS[month]:
                transactions.append(Transaction(date=transaction['date'],user_id=2, balance=transaction['balance'], deposit=transaction['deposit'],
                                    withdrawal=transaction['withdrawal'], description=transaction['description'], investment_total=transaction['investment_total']))
        session.add_all(transactions)
        session.commit()


add_data_db()
