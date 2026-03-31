from app.models import load_model # Import the whole module, not just the variable

def get_bert_prediction(text: str) -> dict:
    """Passes the log through BERT to understand deep context."""
    # Check the module's variable dynamically
    if not load_model.bert_model: 
        return {"prediction": "unknown", "confidence": 0.0}
    
    truncated_text = text[:512] 
    result = load_model.bert_model(truncated_text)[0]
    is_malicious = result['label'] == 'NEGATIVE'
    
    return {
        "prediction": "malicious" if is_malicious else "normal",
        "confidence": round(result['score'], 4)
    }