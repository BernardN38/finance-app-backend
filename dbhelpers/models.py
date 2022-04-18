from sqlalchemy import Column, Float, String,Integer, ForeignKey, Boolean
from sqlalchemy.schema import CheckConstraint
from sqlalchemy.orm import declarative_base, validates
from pprint import pprint

Base = declarative_base()


class Transaction(Base):
    __tablename__ = "transactions"
    id = Column(Integer, primary_key=True)
    user_id = Column(Integer, ForeignKey("users.id"))
    date = Column(String)
    balance = Column(Float)
    deposit = Column(Float)
    withdrawal = Column(Float)
    description = Column(String)
    investment_total = Column(Float)


class User(Base):
    __tablename__ = "users"
    id = Column(Integer, primary_key=True)
    username = Column(String, unique=True, nullable=False)
    first_name = Column(String, nullable=False)
    last_name = Column(String, nullable=False)
    email = Column(String, unique=True, nullable=False)
    password = Column(String, nullable=False)
    is_admin = Column(Boolean, nullable=False, default=False)
    
    __table_args__ = (
        CheckConstraint('char_length(username) > 5',
                        name='username_min_length'),
        CheckConstraint('char_length(first_name) > 5',
                        name='first_name_min_length'),
        CheckConstraint('char_length(last_name) > 5',
                        name='last_name_min_length'),
        CheckConstraint('char_length(email) > 5',
                        name='email_min_length'),
        CheckConstraint('char_length(password) > 5',
                        name='password_min_length'),
    )

    @validates('username')
    def validate_some_string(self, key, some_string) -> str:
        if len(some_string) <= 5:
            raise ValueError('some_string too short')
        return some_string
    @validates('first_name')
    def validate_some_string(self, key, some_string) -> str:
        if len(some_string) <= 5:
            raise ValueError('some_string too short')
        return some_string
    @validates('last_name')
    def validate_some_string(self, key, some_string) -> str:
        if len(some_string) <= 5:
            raise ValueError('some_string too short')
        return some_string
    @validates('email')
    def validate_some_string(self, key, some_string) -> str:
        if len(some_string) <= 5:
            raise ValueError('some_string too short')
        return some_string
    @validates('password')
    def validate_some_string(self, key, some_string) -> str:
        if len(some_string) <= 5:
            raise ValueError('some_string too short')
        return some_string
