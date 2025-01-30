from transformers import AutoTokenizer, AutoModelForTokenClassification
from transformers import pipeline

model_name = "HUMADEX/english_medical_ner"

# Load the tokenizer and model
tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForTokenClassification.from_pretrained(model_name)

# Sample text for inference
text = "I am a 45 year old man who has severe head ache, stomach pain, loose bowel movement and "

# Tokenize the input text
inputs = tokenizer(text, return_tensors="pt")

ner_pipeline = pipeline("ner", model=model, tokenizer=tokenizer)

ner_results = ner_pipeline(text)

for entity in ner_results:
    print(f"{entity['word']} -> {entity['entity']} ({entity['score']:.3f})")
