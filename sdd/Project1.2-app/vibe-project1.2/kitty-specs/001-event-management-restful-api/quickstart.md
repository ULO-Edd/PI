# Quickstart: Event Management RESTful API

This guide provides instructions for setting up and running the FastAPI application.

## Prerequisites

- Python 3.11 or higher
- Pip (Python package installer)

## Setup

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```

2.  **Create and activate a virtual environment:**
    ```bash
    python -m venv venv
    source venv/bin/activate  # On Windows, use `venv\Scripts\activate`
    ```

3.  **Install dependencies:**
    ```bash
    pip install -r requirements.txt
    ```
    *(Note: `requirements.txt` will be created during implementation)*

## Running the Application

1.  **Start the FastAPI server:**
    ```bash
    uvicorn src.api.main:app --reload
    ```

2.  **Access the API documentation:**
    Open your web browser and navigate to `http://127.0.0.1:8000/docs` to view the interactive Swagger UI documentation.

## Running Tests

Execute the following command to run the test suite:

```bash
pytest
```
