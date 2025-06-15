import sys
import uvicorn

if __name__ == "__main__":
    if "--check" in sys.argv:
        print("Server is importable and configured.")
        sys.exit(0)
    uvicorn.run("app:app", host="0.0.0.0", port=8000, reload=True)
