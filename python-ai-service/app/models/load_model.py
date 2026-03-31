from transformers import pipeline
import joblib
import os

# Global variables to hold models in memory
bert_model = None
rf_model = None

def load_all_models():
    global bert_model, rf_model
    print("⏳ Loading AI Models into memory (this may take a moment)...")
    
    # 1. Load BERT
    # For this tutorial, we download a lightweight pre-trained model. 
    # Later, you can point this to "model_store/your-custom-bert"
    try:
        bert_model = pipeline("text-classification", model="distilbert-base-uncased-finetuned-sst-2-english")
        print("✅ BERT Model loaded.")
    except Exception as e:
        print(f"❌ Failed to load BERT: {e}")

    # 2. Load Random Forest
    rf_path = "model_store/rf_model.joblib"
    if os.path.exists(rf_path):
        rf_model = joblib.load(rf_path)
        print("✅ Random Forest Model loaded.")
    else:
        print("⚠️ No RF model found in model_store/. Using fallback heuristics for now.")
        rf_model = None