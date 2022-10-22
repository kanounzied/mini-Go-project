from datetime import date
from http import client
import mysql.connector as connector
from dotenv import load_dotenv
import os
import datetime
import random

SCHEMA_INSERT_FILEPATH = "schema_sql/DB_schema&insertion.sql"

load_dotenv()

cnx = ''

def insertCustomer(client_id, date):
    cursor=cnx.cursor()
    query = "insert into Customer (ClientCustomerID, InsertDate) values (%s, %s)"
    cursor.execute(query, (client_id, date))
    cnx.commit()
    cust_id = cursor.lastrowid
    query = "insert into CustomerData (CustomerID, ChannelTypeID, ChannelValue, InsertDate) values(%s, 1, %s, %s)"
    cursor.execute(query, (
        cust_id,
        'zied' + str(client_id) + '@gmail.com',
        date
    ))
    cnx.commit()
    return cust_id

def insertEvent(client_event_id, date, customer_id):
    cursor=cnx.cursor()
    query = "insert into CustomerEvent (ClientEventID, InsertDate) values (%s, %s)"
    cursor.execute(query, (client_event_id, date))
    cnx.commit()
    event_id = cursor.lastrowid
    query = "insert into CustomerEventData (EventID, ContentID, CustomerID, EventTypeID, EventDate, Quantity) values (%s, %s, %s, %s, %s, %s)"
    cursor.execute(query, (
        event_id,
        random.randint(1, 5),
        customer_id,
        random.randint(1, 6),
        date,
        random.randint(10, 30)
    ))
    cnx.commit()
    return event_id

# create DB schema and insert not user related data
def createAndInsertSchema():
    print("[X] Executing Schema and insert file ... ")
    sql_file = open(SCHEMA_INSERT_FILEPATH)
    sql_as_string = sql_file.read()
    sql_file.close()

    cursor = cnx.cursor()
    cursor.executescript(sql_as_string)
    print("Done ! ")

try:
    cnx = connector.connect(
        user=os.getenv("DBUSER"),
        password=os.getenv("PASSWORD"),
        host=os.getenv("DBURL"),
        database=os.getenv("DBNAME"),
    )
except connector.Error as err:
    print(err)
    print(err.errno)

createAndInsertSchema()

# insert user related data : events
for i in range(100):
        new_customer_id = insertCustomer(i, datetime.datetime.now())
        for event_nb in range(50):
            event_id = insertEvent(event_nb, datetime.datetime.now(), new_customer_id)
print("added ", cnx.cursor().lastrowid, " rows")

cnx.close()