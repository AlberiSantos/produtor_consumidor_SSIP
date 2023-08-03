from flask import Flask, render_template
# import psycopg2
import os
#from dotenv import load_dotenv

app = Flask(__name__)

#load_dotenv()

#DB_NAME = os.getenv("DB_NAME")
#DB_USER = os.getenv("DB_USER")
#DB_PASSWORD = os.getenv("DB_PASSWORD")
#DB_HOST = os.getenv("DB_HOST")

def get_todos():
    try:
        connection = psycopg2.connect(
            dbname=DB_NAME, user=DB_USER, password=DB_PASSWORD, host=DB_HOST
        )
        cursor = connection.cursor()

        cursor.execute("SELECT * FROM tasks;")
        todos = cursor.fetchall()

        return todos

    except Exception as e:
        print("Error connecting to the database:", e)
        return []

    finally:
        cursor.close()
        connection.close()

@app.route('/')
def index():
    # todos = get_todos()
    todos = [{"id": 0, "title": "Sla", "status": False}, {"id": 1, "title": "Sla", "status": False}]
    return render_template('index.html', todos=todos)

if __name__ == '__main__':
    cert_file = os.path.join(os.path.dirname(__file__), '.secrets', 'server.crt')
    key_file = os.path.join(os.path.dirname(__file__), '.secrets', 'server.key')

    context = (cert_file, key_file)
    
    app.run(debug=bool(os.getenv("PRODUCTION")), ssl_context=context)
