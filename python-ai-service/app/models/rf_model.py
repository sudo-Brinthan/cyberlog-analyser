from app.models import load_model

def get_rf_refinement(bert_conf: float, rule_score: int, entity_count: int) -> dict:
    """Uses Random Forest to refine the decision using structured features."""
    
    # Check the module's variable dynamically
    if load_model.rf_model is None:
        # Fallback heuristic
        risk_factor = (bert_conf * 0.4) + (rule_score * 0.15) + (entity_count * 0.05)
        return {
            "prediction": "malicious" if risk_factor > 0.5 else "normal",
            "confidence": min(round(risk_factor, 4), 0.99)
        }