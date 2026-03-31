import re

def clean_log(log_text: str) -> str:
    """Normalizes the log text for the AI models."""
    # Convert to lowercase
    text = log_text.lower()
    # Remove extra whitespaces
    text = re.sub(r'\s+', ' ', text).strip()
    return text