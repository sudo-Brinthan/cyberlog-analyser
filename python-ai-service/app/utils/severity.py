def determine_severity(confidence: float, rule_score: int) -> str:
    """Calculates the final threat level based on AI confidence and rule scores."""
    total_risk = (confidence * 10) + rule_score
    
    if total_risk >= 12:
        return "CRITICAL"
    elif total_risk >= 8:
        return "HIGH"
    elif total_risk >= 4:
        return "MEDIUM"
    else:
        return "LOW"