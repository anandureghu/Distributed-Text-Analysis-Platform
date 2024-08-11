# Distributed Log Processing System

This project is a Distributed Log Processing System built using Go and Docker. It includes multiple microservices designed to handle and process large volumes of log data efficiently.

## **Project Structure**

- **Log Ingestion Service:** Receives log data from external sources and forwards it to the processing service.
- **Log Processing Service:** Processes the incoming log data and distributes it to other services for storage, analysis, or monitoring.
- **Data Storage Service:** Stores the processed log data in a persistent storage system.
- **Analytics Service:** Performs real-time and batch analysis of log data to generate insights.
- **Monitoring Service:** Monitors the entire system, providing metrics and alerts for system health.

## **Getting Started**

### **Prerequisites**

- Go 1.20 or higher
- Docker
- Docker Compose
- [Air](https://github.com/cosmtrek/air) (for hot-reloading during development)

### **Installation**

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/distributed-log-processing-system.git
   cd distributed-log-processing-system
   ```

2. **Build All Docker Images:**

   ```bash
   make build
   ```

3. **Start All Services:**

   ```bash
   make up
   ```

4. **View Logs from All Services:**

   ```bash
   make logs
   ```

5. **Stop All Services:**

   ```bash
   make down
   ```

### **Development with Air**

**Air** is used for hot-reloading during development. You can run Air for each service individually.

1. **Log Ingestion Service:**

   ```bash
   make air-log-ingestion
   ```

2. **Log Processing Service:**

   ```bash
   make air-log-processing
   ```

3. **Data Storage Service:**

   ```bash
   make air-data-storage
   ```

4. **Analytics Service:**

   ```bash
   make air-analytics
   ```

5. **Monitoring Service:**

   ```bash
   make air-monitoring
   ```

## Environment Variables

Configure environment variables in each service by modifying the .env files or directly in the docker-compose.yml file.

## Accessing Services

After starting the services, they will be accessible on the following ports:

Log Ingestion Service: http://localhost:8080  
Log Processing Service: gRPC on port 50051  
Data Storage Service: gRPC on port 50052  
Analytics Service: http://localhost:8081 and gRPC on port 50053  
Monitoring Service: http://localhost:8082

## Contact

For any questions or suggestions, feel free to open an issue or contact me via email.
