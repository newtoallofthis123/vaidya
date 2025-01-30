import torch
from transformers import AutoTokenizer, pipeline
from transformers import BertForTokenClassification

from transformers import AutoTokenizer, AutoModelForTokenClassification

tokenizer = AutoTokenizer.from_pretrained("MilosKosRad/BioNER")
model = AutoModelForTokenClassification.from_pretrained("MilosKosRad/BioNER")
string2 = 'The patient was diagnosed with diabetes, hypertension, asthma, and tuberculosis during the initial health screening'

def ner_prediction(input_string):
    encodings = tokenizer(
        "Disease", 
        input_string,
        is_split_into_words=False,
        padding=True,
        truncation=True,
        add_special_tokens=True,
        return_offsets_mapping=False,
        max_length=512,
        return_tensors="pt"
    )

    # Perform inference
    with torch.no_grad():
        outputs = model(**encodings)

    # Get predicted label IDs and probabilities
    logits = outputs.logits
    predictions = torch.argmax(logits, dim=2)

    # Decode tokens and labels
    tokens = tokenizer.convert_ids_to_tokens(encodings["input_ids"][0])
    labels = [model.config.id2label[p.item()] for p in predictions[0]]

    # Extract entities
    entities = []
    for token, label in zip(tokens, labels):
        if label != "O":  # Ignore tokens labeled as "O" (Outside)
            entities.append({"token": token, "label": label})

    return entities

ner_results = ner_prediction(string2)

for entity in ner_results:
    print(f"{entity['token']} -> {entity['label']}")
