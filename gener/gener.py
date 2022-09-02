import pw
import psycopg2
from psycopg2 import Error
import random
class Query_Handler:
    def __init__(self, user_, password_, host_, port_, database_):
        try:
            self.connection = psycopg2.connect(
                user = user_,
                password = password_,
                host = host_,
                port = port_,
                database = database_)
            self.cursor = self.connection.cursor()
            self.cursor.execute("SELECT version();")
            record = self.cursor.fetchone()
            print("Вы подключены к - ", record, '\n')
            #cursor.close()
            #connection.close()
        except(Exception, Error) as error:
            print("Ошибка при работе с PostgreSQL", error)
    def __del__(self):
        if self.connection:
            self.cursor.close()
            self.connection.close()
            print("CONNECTION CLOSED")
    def my_query(self, query_string):
        try:
            
                self.cursor.execute(query_string)
                try:
                    record = self.cursor.fetchall()
                    for j in record:
                        query_str = " update foundation_tab \
                                      set curFoudrisingAmount = " + str(j[1]) + \
                                    " where id = " + str(j[0])
                        
                        self.cursor.execute(query_str)
                        self.connection.commit()
                        print(query_str)
                except:
                    pass;
        except(Exception, Error) as error:
            print("ERROR IS - ", error)
    def add_country(self):
        cntry = ["USA", "Russia", "UK", "Canada", "France", "Germany", "China", "Italy", "Spain"]
        for i in range(1, 1001):
            query_str = " update foundation_tab \
                                      set country = '" + random.choice(cntry) + "'::text" + \
                                    " where id = " + str(i)
            self.cursor.execute(query_str)
            self.connection.commit()
            print(i)
    def add_username_for_foundation(self):
        for i in range(1, 1001):
            query_str = " select name from  foundation_tab \
                                     where id = " + str(i)
            self.cursor.execute(query_str)
            record = self.cursor.fetchall()
            for j in range (len(record)):
                for k in record[j]:
                    newstr = ""
                    for l in range(len(k)):
                        if k[l] == '\'':
                            continue
                        newstr += k[l]
                    query_str = " update foundation_tab \
                                      set username = '" + newstr + "' " +\
                                    " where id = " + str(i)
                    self.cursor.execute(query_str)
                    self.connection.commit()
                    
    def user_actions(self):
        pass
        '''trans_id = 1
        # foundations
        x = [i for i in range(1, 1001)]
        for i in range (1, 1001):
            summa = random.uniform(1, 10000)
            query_str = " update user_tab \
                                      set balance = balance - " + str(summa) + "::money" + \
                                    " where id = " + str(i)
            self.cursor.execute(query_str)
            self.connection.commit()

            query_str = " update user_tab \
                                      set charitySum = charitySum + " + str(summa) + "::money" + \
                                    " where id = " + str(i)
            self.cursor.execute(query_str)
            self.connection.commit()
            
            foundn_id = x.pop(random.randint(0, len(x) - 1))

            query_str = " update foundation_tab \
                                      set fund_balance = fund_balance + " + str(summa) + "::money" + \
                                    " where id = " + str(foundn_id)
            self.cursor.execute(query_str)
            self.connection.commit()
            

            query_str = " update foundation_tab \
                                      set income_history = income_history + " + str(summa) + "::money" + \
                                    " where id = " + str(foundn_id)
            self.cursor.execute(query_str)
            self.connection.commit()

            query_str = "insert into transaction_tab \
                        (id, from_essence_type, from_id, to_essence_type, sum_of_money, comment, to_id) \
                        values (" + \
                        str(trans_id) + ", " + str(0) + "::bool, " + str(i) + ", " + str(0) + "::bool, " + \
                        str(summa) + "::money, " + "'its donate from user with id = " + str(i) + "', " + str(foundn_id) + ")"
            self.cursor.execute(query_str)
            self.connection.commit()
            trans_id += 1
            print(i)


        # foundrisings
        trans_id = 1001
        x = [i for i in range(1, 1001)]
        for i in range (1, 1001):
            summa = random.uniform(1, 10000)
            query_str = " update user_tab \
                                      set balance = balance - " + str(summa) + "::money" + \
                                    " where id = " + str(i)
            self.cursor.execute(query_str)
            self.connection.commit()

            query_str = " update user_tab \
                                      set charitySum = charitySum + " + str(summa) + "::money" + \
                                    " where id = " + str(i)
            self.cursor.execute(query_str)
            self.connection.commit()
            
            foundn_id = x.pop(random.randint(0, len(x) - 1))

            query_str = " update foundrising_tab \
                                      set current_sum = current_sum + " + str(summa) + "::money" + \
                                    " where id = " + str(foundn_id)
            self.cursor.execute(query_str)
            self.connection.commit()
            
            query_str = "insert into transaction_tab \
                        (id, from_essence_type, from_id, to_essence_type, sum_of_money, comment, to_id) \
                        values (" + \
                        str(trans_id) + ", " + str(0) + "::bool, " + str(i) + ", " + str(1) + "::bool, " + \
                        str(summa) + "::money, " + "'its donate from user with id = " + str(i) + "', " + str(foundn_id) + ")"
            self.cursor.execute(query_str)
            self.connection.commit()
            trans_id += 1
            print(i)'''
            
            
        
    

   

q_ =    "select found_id,count(found_id)\
        from foundrising_tab\
        group by found_id\
        order by found_id"

choice = None
Q1 = Query_Handler (pw.usr, pw.pswd, pw.hst, pw.prt, pw.db)
#Q1.my_query(q_)
Q1.add_username_for_foundation()

