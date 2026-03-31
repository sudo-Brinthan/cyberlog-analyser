from pydantic import BaseModel, Field

class LogRequest(BaseModel):
    # The actual log string from the Go service
    message: str = Field(..., example="Failed password for root from 192.168.1.50 port 22 ssh2")
    
    # Optional metadata you might want to send later
    source: str = Field(default="go-ingestion-service")
    timestamp: str | None = None