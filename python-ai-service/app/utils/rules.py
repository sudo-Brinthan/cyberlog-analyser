def calculate_rule_score(log_text: str) -> int:
    """Applies deterministic security rules to calculate a baseline risk score."""
    score = 0
    text = log_text.lower()
    
    if "failed password" in text or "unauthorized" in text:
        score += 3
    if "root" in text or "admin" in text:
        score += 2
    if "error" in text or "critical" in text:
        score += 1
        
    return min(score, 5) # Cap the rule score at 5