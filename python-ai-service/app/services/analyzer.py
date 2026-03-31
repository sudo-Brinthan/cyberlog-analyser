from app.utils.preprocess import clean_log
from app.utils.ner import extract_entities
from app.utils.rules import calculate_rule_score
from app.utils.severity import determine_severity

# Import our new ML models
from app.models.bert_model import get_bert_prediction
from app.models.rf_model import get_rf_refinement

def run_pipeline(raw_message: str) -> dict:
    # 1. Preprocessing & Extraction
    clean_text = clean_log(raw_message)
    entities = extract_entities(raw_message)
    rule_score = calculate_rule_score(raw_message)
    
    # 2. Deep Learning (BERT Context)
    bert_result = get_bert_prediction(clean_text)
    
    # 3. Traditional ML (Random Forest Refinement)
    rf_result = get_rf_refinement(
        bert_conf=bert_result["confidence"] if bert_result["prediction"] == "malicious" else (1 - bert_result["confidence"]),
        rule_score=rule_score,
        entity_count=len(entities)
    )
    
    # 4. Final Severity Calculation
    severity = determine_severity(rf_result["confidence"], rule_score)
    
    # 5. Output
    return {
        "prediction": rf_result["prediction"],
        "confidence": rf_result["confidence"],
        "entities": entities,
        "rule_score": rule_score,
        "severity_level": severity,
        "bert_base_prediction": bert_result["prediction"] # Included for transparency
    }