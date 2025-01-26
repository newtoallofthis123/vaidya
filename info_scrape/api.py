from flask import Flask, jsonify, request
from flask_cors import CORS

from ml import get_pipe, parse_entities

app = Flask(__name__)
CORS(app)

pipe = get_pipe()

@app.route("/tokenize")
def tokenize():
    text = request.get_json()["text"]

    return jsonify(parse_entities(pipe(text)))

if __name__ == "__main__":
    app.run(debug=True)
