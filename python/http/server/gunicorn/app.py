from flask import Flask

app = Flask(__name__)

@app.route("/")
def response():
	return "r"

if __name__ == "__main__":
	app.run(debug=False)
