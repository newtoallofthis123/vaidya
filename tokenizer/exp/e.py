# Use a pipeline as a high-level helper
from transformers import pipeline

pipe = pipeline("token-classification", model="pruas/BENT-PubMedBERT-NER-Disease")
results = pipe("The patient was diagnosed with diabetes, hypertension, asthma, and tuberculosis during the initial health screening")
print(results)
