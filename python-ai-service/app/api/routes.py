from fastapi import APIRouter
from app.schemas.request_schema import LogRequest
from app.services.analyzer import run_pipeline  # <-- The Brain is imported here

router = APIRouter(prefix="/api/v1", tags=["Threat Intelligence"])

@router.post("/analyze")
async def analyze_log_endpoint(request: LogRequest):
    """
    Receives a raw log from the Go service and passes it to the AI pipeline.
    """
    # 1. Pass the raw log string directly into our new intelligence pipeline
    result = run_pipeline(request.message)
    
    # 2. Attach the original message so the Go service/Dashboard has context
    result["original_message"] = request.message
    
    return result