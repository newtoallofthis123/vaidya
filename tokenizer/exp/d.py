# Use a pipeline as a high-level helper
from transformers import pipeline
pipe = pipeline("token-classification", model="Clinical-AI-Apollo/Medical-NER", aggregation_strategy='simple')
result = pipe('I am a 45 year old woman has pain in the joints and red eyes')

print(result)
