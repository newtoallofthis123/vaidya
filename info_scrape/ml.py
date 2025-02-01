from transformers import pipeline


def get_pipe():
    pipe = pipeline("token-classification", model="ankitcodes/pii_model")
    return pipe


def parse_entities(entity_list):
    """
    Parse list of entity predictions from token classification into key-value pairs.

    Args:
        entity_list (list): List of dictionaries containing token classification results
            with 'entity' and 'word' keys for each token

    Returns:
        dict: Dictionary mapping entity types to their combined token values
              with word piece tokens properly joined
    """
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
