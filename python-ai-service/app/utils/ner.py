import re

def extract_entities(log_text: str) -> list[str]:
    """Extracts critical entities like IP addresses from the log."""
    entities = []
    # Simple regex to find IPv4 addresses
    ip_pattern = r'\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b'
    ips = re.findall(ip_pattern, log_text)
    entities.extend(ips)
    return entities