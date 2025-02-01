from transformers import AutoTokenizer, AutoModelForTokenClassification
from transformers import pipeline


class SymptomTokenizer:
    """
    Initializes the SymptomTokenizer with a pre-trained model and tokenizer.
    """

    def __init__(self) -> None:
        self.model_name = "HUMADEX/english_medical_ner"
        print("Hello World")

        self.tokenizer = AutoTokenizer.from_pretrained(self.model_name)
        self.model = AutoModelForTokenClassification.from_pretrained(self.model_name)

        self.pipeline = pipeline("ner", model=self.model, tokenizer=self.tokenizer)

    def group_entities(self, tokens):
        """
        Groups tokens into entities based on their tags.

        Args:
            tokens (list): List of tokens with their respective tags.

        Returns:
            list: List of grouped entities.

        Raises:
            ValueError: If there are inconsistencies in the token tags.
        """
        problems = []
        cur = None

        for token in tokens:
            tag = token["entity"]

            if tag.startswith("B"):
                if cur is not None:
                    cur.append(token)
                    problems.append(cur)
                    cur = None
                cur = [token]
            elif tag.startswith("I"):
                if cur is None:
                    raise ValueError("Found I-PROBLEM without preceding B-PROBLEM")
                cur.append(token)
            elif tag.startswith("E") or tag.startswith("S"):
                if cur is None:
                    cur = []
                cur.append(token)
                problems.append(cur)
                cur = None

        if cur is not None:
            raise ValueError("Unclosed B-PROBLEM at end of token list")

        return problems

    def get_symptoms(self, text: str) -> dict[str, list[dict[str, str]]]:
        """
        Extracts symptoms, treatments, and tests from the given text.

        Args:
            text (str): Input text to be tokenized and analyzed.

        Returns:
            dict: Dictionary containing lists of symptoms, treatments, and tests.
        """
        tokens = self.tokenize_text(text)

        print(tokens)

        entites = self.group_entities(tokens)

        res = {"symptoms": [], "treatments": [], "tests": []}
        for group in entites:
            token = {
                "name": "",
                "loc": "",
                "confidences": [],
            }
            for i, e in enumerate(group):
                word = e["word"]
                if not word.startswith("#"):
                    word += " "
                else:
                    word = word.replace("#", "")
                token["name"] += word
                if i == 0:
                    token["loc"] += str(e["start"]) + ","
                if i == len(group) - 1:
                    token["loc"] += str(e["end"])
                token["confidences"].append(e["score"])
            token["name"] = token["name"][:-1]
            token["score"] = sum(token["confidences"]) / len(token["confidences"])
            if group[0]["entity"].endswith("PROBLEM"):
                res["symptoms"].append(token)
            elif group[0]["entity"].endswith("TREATMENT"):
                res["treatments"].append(token)
            elif group[0]["entity"].endswith("TEST"):
                res["tests"].append(token)

        return res

    def tokenize_text(self, text: str):
        """
        Tokenizes the input text using the pre-trained pipeline.

        Args:
            text (str): Input text to be tokenized.

        Returns:
            list: List of tokens with their respective tags.
        """
        return self.pipeline(text)
