# Use a pipeline as a high-level helper
from transformers import pipeline

pipe = pipeline("token-classification", model="ankitcodes/pii_model")
results = pipe("Hello! I am Ishan Joshi, a student from India. I currently live in Mayuri Nagar, Hyderbad. I was born on 9th december 2004. My phone number is +919014678452 and email is noobscience123@gmail.com")

# for res in results:
#     print(f'{res["entity"]}: {res["word"]}')

def parse_entities(entity_list):
    parsed_data = {}
    current_entity = None
    current_value = []

    for item in entity_list:
        tag, word = item["entity"].lower(), item["word"]

        if tag.startswith("b-"):
            if current_entity:
                parsed_data[current_entity] = "".join(current_value).replace("##", "")
            current_entity = tag[2:]
            current_value = [word]
        elif tag.startswith("i-") and current_entity == tag[2:]:
            current_value.append(word)
        elif tag.startswith("e-") and current_entity == tag[2:]:
            current_value.append(word)
            parsed_data[current_entity] = "".join(current_value).replace("##", "")
            current_entity = None
            current_value = []
        else:
            current_entity = None
            current_value = []

    if current_entity:
        parsed_data[current_entity] = "".join(current_value).replace("##", "")

    return parsed_data

print(parse_entities(results))
