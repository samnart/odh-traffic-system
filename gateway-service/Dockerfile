# Use a specific OpenJDK version to avoid security vulnerabilities and for better reproducibility
FROM eclipse-temurin:17-jre-alpine

# Set the working directory
WORKDIR /app

# Add a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Create directory for the application
RUN mkdir -p /app/logs && chown -R appuser:appgroup /app

# Set environment variables
ENV JAVA_OPTS="-Xms256m -Xmx512m -XX:+UseG1GC -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/app/logs"
ENV SERVER_PORT=8080

# Copy the JAR file
COPY target/*.jar /app/app.jar

# Change ownership of the application files
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose the application port
EXPOSE ${SERVER_PORT}

# Set healthcheck
HEALTHCHECK --interval=30s --timeout=5s --start-period=30s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:${SERVER_PORT}/actuator/health || exit 1

# Run the application
ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -jar /app/app.jar"]