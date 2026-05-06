# # Metamorphosis Retrieval Augmented Generation
Basic RAG for the book `Metamorphosis` by Franz Kafka

### Models Used:
- LLM - **Mistral:7b-instruct-q4_K_M** since this is the most sensible choice for a 4060 gpu
- Embedding Model - **nomic-embed-text** because its lightweight and works well for retrieval tasks

### Vector Store
- Qdrant - it has Go client and has fast vector search and filtering